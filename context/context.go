package context

import "github.com/PeerXu/jarvis3/user"

type Context interface {
	User() *user.User
	AccessToken() *user.AccessToken
}

type jarvisContext struct {
	user        *user.User
	accessToken *user.AccessToken
}

func (ctx jarvisContext) User() *user.User {
	return ctx.user
}

func (ctx jarvisContext) AccessToken() *user.AccessToken {
	return ctx.accessToken
}

func NewContext(u *user.User, t *user.AccessToken) Context {
	return jarvisContext{
		user:        u,
		accessToken: t,
	}
}

func MustNewContextByUser(u *user.User) Context {
	return jarvisContext{
		user:        u,
		accessToken: u.AccessTokens[0],
	}
}
