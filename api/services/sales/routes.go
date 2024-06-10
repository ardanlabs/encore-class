package sales

import (
	"context"
	"net/http"

	"encore.dev"
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
//encore:api public method=POST path=/test
func (s *Service) Test(ctx context.Context, status testapp.Status) (testapp.Status, error) {
	return testapp.Test(ctx, status)
}
