package mid

import (
	"fmt"
	"runtime/debug"

	"encore.dev/beta/errs"
	"encore.dev/middleware"
	"github.com/ardanlabs/encore/app/sdk/metrics"
)

// Panics handles panics that occur when processing a request.
func Panics(v *metrics.Values, req middleware.Request, next middleware.Next) (resp middleware.Response) {
	defer func() {
		if rec := recover(); rec != nil {
			trace := debug.Stack()

			resp = middleware.Response{
				Err: &errs.Error{
					Code:    errs.Internal,
					Message: fmt.Sprintf("PANIC [%v] TRACE[%s]", rec, string(trace)),
				},
			}

			v.IncPanics()
		}
	}()

	return next(req)
}
