// Package mid provides context support.
package mid

import (
	"context"
	"errors"

	"encore.dev/middleware"
	"github.com/ardanlabs/encore/app/sdk/auth"
	"github.com/ardanlabs/encore/business/domain/userbus"
	"github.com/google/uuid"
)

// AuthInfo defines the information required to perform an authorization.
type AuthInfo struct {
	Claims auth.Claims
	UserID uuid.UUID
	Rule   string
}

// =============================================================================

type ctxKey int

const (
	userKey ctxKey = iota + 1
)

func setUser(req middleware.Request, usr userbus.User) middleware.Request {
	ctx := context.WithValue(req.Context(), userKey, usr)
	return req.WithContext(ctx)
}

// GetUser extracts the user from the context.
func GetUser(ctx context.Context) (userbus.User, error) {
	v, ok := ctx.Value(userKey).(userbus.User)
	if !ok {
		return userbus.User{}, errors.New("user not found")
	}

	return v, nil
}
