package sales

import (
	"context"

	"github.com/ardanlabs/encore/app/domain/testapp"
)

// =============================================================================

//lint:ignore U1000 "called by encore"
//encore:api public method=POST path=/test
func (s *Service) Test(ctx context.Context, status testapp.Status) (testapp.Status, error) {
	return testapp.Test(ctx, status)
}
