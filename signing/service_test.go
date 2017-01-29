package signing

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	c "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	"github.com/PeerXu/jarvis3/repository"
)

func TestSigningService(t *testing.T) {

	c.Convey("singing service", t, func() {
		users := repository.NewUserRepository()
		logger := log.NewLogfmtLogger(os.Stderr)
		s := NewService(logger, users)
		ctx := context.Background()

		c.Convey("login error password", func() {
			_, err := s.Login(ctx, "admin", "??????")
			c.So(err, c.ShouldNotBeNil)
		})

		c.Convey("login an existed user", func() {
			u, err := s.Login(ctx, "admin", "admin")
			c.So(err, c.ShouldBeNil)

			at := u.AccessTokens[0]

			jctx := jcontext.NewContext(u, at)
			ctx = context.WithValue(ctx, "JarvisContext", jctx)

			c.Convey("create a new user", func() {
				u2, err := s.CreateUser(ctx, "test", "test", "test@jarvis3.cc")
				c.So(err, c.ShouldBeNil)

				c.Convey("login with new user", func() {
					ctx = context.Background()
					_, err = s.Login(ctx, "test", "test")
					c.So(err, c.ShouldBeNil)
				})

				c.Convey("delete a new user", func() {
					err = s.DeleteUserByID(ctx, u2.ID)
					c.So(err, c.ShouldBeNil)

					_, err = s.GetUserByID(ctx, u2.ID)
					c.So(err, c.ShouldNotBeNil)
				})
			})

			c.Convey("create an agent token", func() {
				_, err := s.CreateAgentToken(ctx, "token1")
				c.So(err, c.ShouldBeNil)

				c.Convey("delete an agent token", func() {
					err = s.DeleteAgentToken(ctx, "token1")
					c.So(err, c.ShouldBeNil)
				})
			})

			c.Convey("validate access token", func() {
				ctx = context.WithValue(ctx, "Authorization", at.Token)
				_, err = s.ValidateToken(ctx)
				c.So(err, c.ShouldBeNil)
			})

			c.Convey("logout with access-token", func() {
				err = s.Logout(ctx, u.ID)
				c.So(err, c.ShouldBeNil)
			})
		})
	})
}
