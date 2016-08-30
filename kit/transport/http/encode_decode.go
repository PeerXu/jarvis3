package http

import (
	"net/http"

	"golang.org/x/net/context"
)

func ExtractAuthorizationFromRequestHeader(ctx context.Context, r *http.Request) context.Context {
	authorization := r.Header.Get("Authorization")
	ctx = context.WithValue(ctx, "Authorization", authorization)
	return ctx
}
