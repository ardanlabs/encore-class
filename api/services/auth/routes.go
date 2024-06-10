package auth

import "context"

type token struct {
	Value string
}

//lint:ignore U1000 "called by encore"
//encore:api public method=GET path=/token/:kid
func Token(ctx context.Context, kid string) (token, error) {
	t := token{
		Value: kid + "-1",
	}

	return t, nil
}