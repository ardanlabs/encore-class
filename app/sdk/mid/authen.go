package mid

import (
	"context"
	"fmt"

	eauth "encore.dev/beta/auth"
	"github.com/ardanlabs/encore/app/sdk/auth"
	"github.com/ardanlabs/encore/app/sdk/errs"
	"github.com/google/uuid"
)

// Bearer processes JWT authentication logic.
func Bearer(ctx context.Context, ath *auth.Auth, authorization string) (eauth.UID, *auth.Claims, error) {
	claims, err := ath.Authenticate(ctx, authorization)
	if err != nil {
		return "", nil, errs.New(errs.Unauthenticated, err)
	}

	if claims.Subject == "" {
		return "", nil, errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, no claims")
	}

	subjectID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return "", nil, errs.New(errs.Unauthenticated, fmt.Errorf("parsing subject: %w", err))
	}

	return eauth.UID(subjectID.String()), &claims, nil
}
