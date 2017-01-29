package repository

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"

	"github.com/PeerXu/jarvis3/user"
)

func TestUserRepository(t *testing.T) {
	c.Convey("Testing User Repository", t, func() {
		r := NewUserRepository()

		c.Convey("should create a user", func() {
			u := user.NewUser("admin", "admin", "jarvis3@gmail.com")
			_, err := r.CreateUser(u)
			c.So(err, c.ShouldBeNil)

			userID := u.ID

			_, err = r.GetUserByID(userID)
			c.So(err, c.ShouldBeNil)

			c.Convey("should create an access-token", func() {
				t := user.NewAccessToken()
				err = r.CreateAccessToken(userID, t)
				c.So(err, c.ShouldBeNil)

				u2, err := r.LookupUserByAccessToken(t)
				c.So(err, c.ShouldBeNil)
				c.So(u2.Username, c.ShouldEqual, u.Username)

				c.Convey("should delete an existed access-token", func() {
					err = r.DeleteAccessTokens(userID, []*user.AccessToken{t})
					c.So(err, c.ShouldBeNil)

					_, err := r.LookupUserByAccessToken(t)
					c.So(err, c.ShouldEqual, user.ErrAccessTokenNotFound)
				})

			})

			c.Convey("should create an agent-token", func() {
				t := user.NewAgentToken("app")
				err = r.CreateAgentToken(userID, t)
				c.So(err, c.ShouldBeNil)

				u2, err := r.LookupUserByAgentToken(t)
				c.So(err, c.ShouldBeNil)
				c.So(u2.Username, c.ShouldEqual, u.Username)

				t2, err := r.LookupAgentTokenByName(userID, "app")
				c.So(err, c.ShouldBeNil)
				c.So(t2.Token, c.ShouldEqual, t.Token)

				c.Convey("should delete an existed agent-token", func() {
					err = r.DeleteAgentTokens(userID, []*user.AgentToken{t})
					c.So(err, c.ShouldBeNil)

					_, err := r.LookupUserByAgentToken(t)
					c.So(err, c.ShouldEqual, user.ErrAgentTokenNotFound)

					_, err = r.LookupAgentTokenByName(userID, "app")
					c.So(err, c.ShouldEqual, user.ErrAgentTokenNotFound)
				})
			})

			c.Convey("should delete an existed user", func() {
				err := r.DeleteUserByID(userID)
				c.So(err, c.ShouldBeNil)
				_, err = r.GetUserByID(userID)
				c.So(err, c.ShouldEqual, user.ErrUserNotFound)
			})
		})

	})
}
