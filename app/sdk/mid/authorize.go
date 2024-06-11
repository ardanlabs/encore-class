package mid

import (
	"errors"

	eauth "encore.dev/beta/auth"
	"encore.dev/middleware"
	"github.com/ardanlabs/encore/app/sdk/auth"
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
