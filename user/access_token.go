package user

import (
	"strings"
	"time"
)

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

func ParseAccessTokenFromString(s string) (*AccessToken, error) {
	xs := strings.Split(s, ":")
	if len(xs) != 2 {
		return nil, ErrAccessTokenNotFound
	}

	typ, token := xs[0], xs[1]

	if strings.ToLower(strings.Trim(typ, " ")) != "act" {
		return nil, ErrAccessTokenNotFound
	}

	return &AccessToken{Token: strings.Trim(token, " ")}, nil
}
