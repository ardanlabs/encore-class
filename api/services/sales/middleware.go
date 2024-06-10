package sales

import (
	"encore.dev/middleware"
	"github.com/ardanlabs/encore/app/sdk/mid"
)

// NOTE: The order matters so be careful when injecting new middleware. Global
//       middleware will always come first. We want the Auth middleware to
//       happen before any non-global middlware.

// =============================================================================
// Global middleware functions

//lint:ignore U1000 "called by encore"
//encore:middleware target=all
func (s *Service) panics(req middleware.Request, next middleware.Next) middleware.Response {
	return mid.Panics(s.mtrcs, req, next)
}

// =============================================================================
// Specific middleware functions

//lint:ignore U1000 "called by encore"
//encore:middleware target=tag:metrics
func (s *Service) metrics(req middleware.Request, next middleware.Next) middleware.Response {
	return mid.Metrics(s.mtrcs, req, next)
}
