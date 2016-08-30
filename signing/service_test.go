package signing

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	c "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	"github.com/PeerXu/jarvis3/repository"
	"github.com/PeerXu/jarvis3/user"
)

func TestSigningService(t *testing.T) {
	users := repository.NewUserRepository()
	logger := log.NewLogfmtLogger(os.Stderr)
	s := NewService(logger, users)
	ctx := context.Background()

	c.Convey("login error password", t, func() {
		_, err := s.Login(ctx, "admin", "??????")
		c.So(err, c.ShouldNotBeNil)
	})

	c.Convey("login an existed user", t, func() {
		at, err := s.Login(ctx, "admin", "admin")
		c.So(err, c.ShouldBeNil)

		jctx := jcontext.NewContext(&user.User{Username: "admin"}, at)
		ctx = context.WithValue(ctx, "JarvisContext", jctx)

		c.Convey("create a new user", func() {
			_, err := s.CreateUser(ctx, "test", "test", "test@jarvis3.cc")
			c.So(err, c.ShouldBeNil)

			c.Convey("login with new user", func() {
				ctx = context.Background()
				at, err = s.Login(ctx, "test", "test")
				c.So(err, c.ShouldBeNil)
			})

			c.Convey("delete a new user", func() {
				err = s.DeleteUser(ctx, "test")
				c.So(err, c.ShouldBeNil)

				_, err = s.GetUser(ctx, "test")
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

		c.Convey("logout with access-token", func() {
			err = s.Logout(ctx, "admin")
			c.So(err, c.ShouldBeNil)
		})

		c.Convey("validate access token", func() {
			ctx = context.WithValue(ctx, "Authorization", at.Token)
			_, err = s.ValidateAccessToken(ctx)
			c.So(err, c.ShouldBeNil)
		})
	})
}
