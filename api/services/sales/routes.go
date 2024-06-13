package sales

import (
	"context"
	"net/http"

	"encore.dev"
	"encore.dev/beta/errs"
	"github.com/ardanlabs/encore/app/domain/testapp"
	"github.com/ardanlabs/encore/app/domain/userapp"
	"github.com/ardanlabs/encore/app/sdk/query"
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

// =============================================================================

//lint:ignore U1000 "called by encore"
//encore:api auth method=POST path=/v1/users tag:metrics tag:authorize tag:as_admin_role
func (s *Service) UserCreate(ctx context.Context, app userapp.NewUser) (userapp.User, error) {
	return s.userApp.Create(ctx, app)
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=PUT path=/v1/users/:userID tag:metrics tag:authorize_user
func (s *Service) UserUpdate(ctx context.Context, userID string, app userapp.UpdateUser) (userapp.User, error) {
	return s.userApp.Update(ctx, app)
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=PUT path=/v1/role/:userID tag:metrics tag:authorize_user tag:as_admin_role
func (s *Service) UserUpdateRole(ctx context.Context, userID string, app userapp.UpdateUserRole) (userapp.User, error) {
	return s.userApp.UpdateRole(ctx, app)
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=DELETE path=/v1/users/:userID tag:metrics tag:authorize_user
func (s *Service) UserDelete(ctx context.Context, userID string) error {
	return s.userApp.Delete(ctx)
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=GET path=/v1/users tag:metrics tag:authorize tag:as_admin_role
func (s *Service) UserQuery(ctx context.Context, qp userapp.QueryParams) (query.Result[userapp.User], error) {
	return s.userApp.Query(ctx, qp)
}

//lint:ignore U1000 "called by encore"
//encore:api auth method=GET path=/v1/users/:userID tag:metrics tag:authorize_user
func (s *Service) UserQueryByID(ctx context.Context, userID string) (userapp.User, error) {
	return s.userApp.QueryByID(ctx)
}
