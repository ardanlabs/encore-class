// Package testapp was an example.
package testapp

import "context"

func Test(ctx context.Context, status Status) (Status, error) {
	return status, nil
}
