package http

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	jerrors "github.com/PeerXu/jarvis3/errors"
	"github.com/PeerXu/jarvis3/utils"
)

func DecodeErrorResponseMiddleware(next httptransport.DecodeResponseFunc) httptransport.DecodeResponseFunc {
	return func(ctx context.Context, res *http.Response) (interface{}, error) {
		if res.StatusCode != http.StatusNoContent {
			body, err := utils.ReadAndAssignResponseBody(res)
			if err != nil {
				return nil, err
			}

			jerr, ok, err := jerrors.DecodeJarvisError(body)
			if err != nil {
				return nil, err
			}
			if ok {
				return jerr, nil
			}
		}

		return next(ctx, res)
	}
}

func LoadEnvironmentIntoContextMiddleware(e utils.Environment) func(httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
	return func(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
		return func(ctx context.Context, r *http.Request, request interface{}) error {
			ctx = context.WithValue(ctx, "AccessToken", e.Get("AccessToken"))
			ctx = context.WithValue(ctx, "AgentToken", e.Get("AgentToken"))
			ctx = context.WithValue(ctx, "UserID", e.Get("UserID"))
			return next(ctx, r, request)
		}
	}
}

func LoadRequestIntoContextMiddleware() func(httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
	return func(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
		return func(ctx context.Context, r *http.Request, request interface{}) error {
			ctx = context.WithValue(ctx, "AccessToken", r.Header.Get("JVS_ACCESS_TOKEN"))
			return next(ctx, r, request)
		}
	}
}

func AssignAccessTokenToContextMiddleware0(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
	return func(ctx context.Context, r *http.Request, request interface{}) error {
		ati := ctx.Value("Authorization")
		if ati == nil || ati.(string) == "" {
			ati = ctx.Value("AccessToken")
		}
		if ati != nil {
			at := ati.(string)
			r.Header.Set("Authorization", at)
		}
		return next(ctx, r, request)
	}
}

func AssignTokenToContextMiddleware(prefix, key string) func(httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
	return func(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
		return func(ctx context.Context, r *http.Request, request interface{}) error {
			ati := ctx.Value("Authorization")
			if ati == nil || ati.(string) == "" {
				ati = ctx.Value(key)
			}
			if ati != nil {
				at := ati.(string)
				r.Header.Set("Authorization", prefix+at)
			}
			return next(ctx, r, request)
		}
	}
}

var AssignAccessTokenToContextForClientMiddleware = AssignTokenToContextMiddleware("ACT: ", "AccessToken")
var AssignAccessTokenToContextForServiceMiddleware = AssignTokenToContextMiddleware("", "AccessToken")
var AssignAgentTokenToContextForClientMiddleware = AssignTokenToContextMiddleware("AGT: ", "AgentToken")
var AssignAgentTokenToContextForServiceMiddleware = AssignTokenToContextMiddleware("", "AgentToken")
