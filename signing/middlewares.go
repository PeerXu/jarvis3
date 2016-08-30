package signing

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	"github.com/PeerXu/jarvis3/user"
	"github.com/PeerXu/jarvis3/utils"
)

func ValidateAccessTokenMiddeware(s Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			u, err := s.ValidateAccessToken(ctx)
			if err != nil {
				switch err {
				case user.ErrInvalidUsernameOrPassword:
					err = newSignError(errorAccessDenied, "access denied", err)
				}
				return err, nil
			}

			tokenStr := ctx.Value("Authorization").(string)
			var token *user.AccessToken

			for _, t := range u.AccessTokens {
				if t.Token == tokenStr {
					token = t
					break
				}
			}

			jctx := jcontext.NewContext(u, token)
			ctx = context.WithValue(ctx, "JarvisContext", jctx)

			return next(ctx, request)
		}
	}
}

func DecodeErrorResponseMiddleware(next httptransport.DecodeResponseFunc) httptransport.DecodeResponseFunc {
	return func(ctx context.Context, res *http.Response) (interface{}, error) {
		if res.StatusCode != http.StatusNoContent {
			body, err := utils.ReadAndAssignResponseBody(res)
			if err != nil {
				return nil, err
			}

			var serr signError
			err = json.NewDecoder(body).Decode(&serr)
			if err != nil {
				return nil, err
			}

			if serr.Type != "" {
				return &serr, nil
			}
		}

		return next(ctx, res)
	}
}

func LoadEnvironmentIntoContextMiddleware(e utils.Environment) func(httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
	return func(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
		return func(ctx context.Context, r *http.Request, request interface{}) error {
			ctx = context.WithValue(ctx, "AccessToken", e.Get("AccessToken"))
			ctx = context.WithValue(ctx, "Username", e.Get("Username"))
			return next(ctx, r, request)
		}
	}
}

func AssignAccessTokenToContextMiddleware(next httptransport.EncodeRequestFunc) httptransport.EncodeRequestFunc {
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
