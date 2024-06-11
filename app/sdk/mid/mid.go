// Package mid provides context support.
package mid

import (
	"github.com/ardanlabs/encore/app/sdk/auth"
	"github.com/google/uuid"
)

// AuthInfo defines the information required to perform an authorization.
type AuthInfo struct {
	Claims auth.Claims
	UserID uuid.UUID
	Rule   string
}
