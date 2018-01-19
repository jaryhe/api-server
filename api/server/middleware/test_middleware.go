package middleware

import (
	"fmt"
	"net/http"
	"golang.org/x/net/context"
)
type TestMiddleware struct {
	test string
}

func NewTestMiddleware(testString string) TestMiddleware{
	return TestMiddleware{test:testString,}
}

func (v TestMiddleware) WrapHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Println(v.test)
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		return handler(ctx, w, r, vars)
	}
}
