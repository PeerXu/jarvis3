package computing

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	signing_service "github.com/PeerXu/jarvis3/signing/service"
)

func RemoteValidateTokenMiddleware(cli signing_service.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			u, err := cli.ValidateToken(ctx)
			if err != nil {
				return err, nil
			}

			jctx := jcontext.MustNewContextByUser(u)
			ctx = context.WithValue(ctx, "JarvisContext", jctx)

			return next(ctx, request)
		}
	}
}
