package user

import "time"

type AccessToken struct {
	Token    string
	ExpireAt time.Time
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		Token:    genTokenString(),
		ExpireAt: time.Now().Add(30 * time.Minute),
	}
}
