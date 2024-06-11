package sales

import (
	"context"
	"net/http"

	"encore.dev"
	"encore.dev/beta/errs"
	"github.com/ardanlabs/encore/app/domain/testapp"
)

// Fallback is called for the debug enpoints.
//
//encore:api public raw path=/!fallback
func (s *Service) Fallback(w http.ResponseWriter, r *http.Request) {

	// If this is a web socket call for statsviz and we are in development.
	if r.URL.String() == "/debug/statsviz/ws" && encore.Meta().Environment.Type == encore.EnvDevelopment {

		// In development the r.Host will be host=127.0.0.1:RandPort while the
		// Origin will be origin=127.0.0.1:4000. These need to match.
		r.Header.Set("Origin", "http://"+r.Host)
	}

	s.debug.ServeHTTP(w, r)
}

// =============================================================================

//lint:ignore U1000 "called by encore"
//encore:api auth method=POST path=/test tag:metrics tag:authorize tag:as_user_role
func (s *Service) Test(ctx context.Context, status testapp.Status) (testapp.Status, error) {
	return testapp.Test(ctx, status)
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=POST path=/testerror tag:metrics tag:authorize tag:as_user_role
func (s *Service) TestError(ctx context.Context, status testapp.Status) (testapp.Status, error) {
	err := &errs.Error{
		Code:    errs.FailedPrecondition,
		Message: "missing something",
	}

	return testapp.Status{}, err
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=POST path=/testpanic tag:metrics tag:authorize tag:as_user_role
func (s *Service) TestPanic(ctx context.Context, status testapp.Status) (testapp.Status, error) {
	panic("THIS IS A PANIC")

	return testapp.Status{}, nil
}
