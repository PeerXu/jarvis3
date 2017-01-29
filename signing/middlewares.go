package signing

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	jerrors "github.com/PeerXu/jarvis3/errors"
	. "github.com/PeerXu/jarvis3/signing/error"
	. "github.com/PeerXu/jarvis3/signing/service"
	"github.com/PeerXu/jarvis3/user"
)

func ValidateTokenMiddeware(s Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			u, err := s.ValidateToken(ctx)
			if err != nil {
				switch err {
				case user.ErrInvalidUsernameOrPassword:
					err = NewSignError(jerrors.ErrorAccessDenied, "access denied", err)
				}
				return err, nil
			}

			tokenStr := ctx.Value("Authorization").(string)
			var ok bool
			var token *user.AccessToken

			if token, ok = u.LookupAccessTokenByString(tokenStr); !ok {
				return NewSignError(jerrors.ErrorAccessDenied, "access denied", nil), nil
			}

			jctx := jcontext.NewContext(u, token)
			ctx = context.WithValue(ctx, "JarvisContext", jctx)

			return next(ctx, request)
		}
	}
}
