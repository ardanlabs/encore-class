package mid

import (
	"errors"
	"fmt"

	eauth "encore.dev/beta/auth"
	"encore.dev/middleware"
	"github.com/ardanlabs/encore/app/sdk/auth"
	"github.com/ardanlabs/encore/business/domain/userbus"
	"github.com/google/uuid"
)

// ErrInvalidID represents a condition where the id is not a uuid.
var ErrInvalidID = errors.New("ID is not in its proper form")

// Authorize checks the user making the request is an admin or user.
func Authorize(req middleware.Request) (AuthInfo, middleware.Request, error) {
	claims := eauth.Data().(*auth.Claims)

	rule := auth.RuleAdminOnly
	for _, tag := range req.Data().API.Tags {
		switch tag {
		case "as_any_role":
			rule = auth.RuleAny
		case "as_user_role":
			rule = auth.RuleUserOnly
		}
	}

	authInfo := AuthInfo{
		Claims: *claims,
		UserID: uuid.UUID{},
		Rule:   rule,
	}

	// We should call the Auth Service from here and keep things in the app
	// layer but Encore won't allow it. The API layer middleware calls
	// this function first and then calls the Auth Service. This is the same
	// for the other Authorize middleware functions.

	return authInfo, req, nil
}

// AuthorizeUser checks the user making the call has specified a user id on
// the route that matches the claims.
func AuthorizeUser(userBus *userbus.Business, req middleware.Request) (AuthInfo, middleware.Request, error) {
	ctx := req.Context()
	var userID uuid.UUID

	rule := auth.RuleAdminOrSubject
	for _, tag := range req.Data().API.Tags {
		if tag == "as_admin_role" {
			rule = auth.RuleAdminOnly
			break
		}
	}

	if len(req.Data().PathParams) == 1 {
		id := req.Data().PathParams[0]

		var err error
		userID, err = uuid.Parse(id.Value)
		if err != nil {
			return AuthInfo{}, req, ErrInvalidID
		}

		usr, err := userBus.QueryByID(ctx, userID)
		if err != nil {
			switch {
			case errors.Is(err, userbus.ErrNotFound):
				return AuthInfo{}, req, err

			default:
				return AuthInfo{}, req, fmt.Errorf("querybyid: userID[%s]: %s", userID, err)
			}
		}

		req = setUser(req, usr)
	}

	claims := eauth.Data().(*auth.Claims)

	authInfo := AuthInfo{
		Claims: *claims,
		UserID: userID,
		Rule:   rule,
	}

	return authInfo, req, nil
}
