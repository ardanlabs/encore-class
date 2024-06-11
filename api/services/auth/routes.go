package auth

import (
	"context"
	"strings"

	eauth "encore.dev/beta/auth"
	"github.com/ardanlabs/encore/app/sdk/auth"
	"github.com/ardanlabs/encore/app/sdk/errs"
	"github.com/ardanlabs/encore/app/sdk/mid"
)

// =============================================================================
// JWT or Basic Athentication handling

type authParams struct {
	Authorization string `header:"Authorization"`
}

//lint:ignore U1000 "called by encore"
//encore:authhandler
func (s *Service) AuthHandler(ctx context.Context, ap *authParams) (eauth.UID, *auth.Claims, error) {
	parts := strings.Split(ap.Authorization, " ")

	switch parts[0] {
	case "Bearer":
		return mid.Bearer(ctx, s.auth, ap.Authorization)
	}

	return "", nil, errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action")
}

//lint:ignore U1000 "called by encore"
//encore:api private method=POST path=/authorize
func (s *Service) Authorize(ctx context.Context, authInfo mid.AuthInfo) error {
	if err := s.auth.Authorize(ctx, authInfo.Claims, authInfo.UserID, authInfo.Rule); err != nil {
		return errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, claims[%v] rule[%v]: %s", authInfo.Claims.Roles, authInfo.Rule, err)
	}

	return nil
}
