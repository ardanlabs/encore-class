package sales

import (
	"context"
	"fmt"
	"time"

	"encore.dev/middleware"
	authsrv "github.com/ardanlabs/encore/api/services/auth"
	"github.com/ardanlabs/encore/app/sdk/errs"
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
// Authorization related middleware

//lint:ignore U1000 "called by encore"
//encore:middleware target=tag:authorize
func (s *Service) authorize(req middleware.Request, next middleware.Next) middleware.Response {
	p, req, err := mid.Authorize(req)
	if err != nil {
		return errs.NewResponse(errs.Unauthenticated, err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	if err := authsrv.Authorize(ctx, p); err != nil {
		err = fmt.Errorf("%s", err.Error()[17:]) // Remove "unauthenticated:" from the error string.
		return errs.NewResponse(errs.Unauthenticated, err)
	}

	return next(req)
}

// =============================================================================
// Specific middleware functions

//lint:ignore U1000 "called by encore"
//encore:middleware target=tag:metrics
func (s *Service) metrics(req middleware.Request, next middleware.Next) middleware.Response {
	return mid.Metrics(s.mtrcs, req, next)
}
