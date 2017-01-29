package user

import (
	"time"

	"github.com/PeerXu/jarvis3/utils"
)

type UserID string

func (id UserID) String() string {
	return string(id)
}

type User struct {
	ID       UserID
	Username string
	Password EncryptedPassword
	Email    Email

	AccessTokens []*AccessToken
	AgentTokens  []*AgentToken
}

func NewUser(username, password, email string) *User {
	return &User{
		ID:           UserID(utils.NewRandomUUIDString()),
		Username:     username,
		Password:     NewEncryptedPassword(password),
		Email:        Email(email),
		AccessTokens: []*AccessToken{},
		AgentTokens:  []*AgentToken{},
	}
}

func (user *User) ValidatePassword(password string) bool {
	return user.Password.Validate(password)
}

func (user *User) LookupAccessTokenByString(s string) (*AccessToken, bool) {
	now := time.Now()
	token, err := ParseAccessTokenFromString(s)
	if err != nil {
		return nil, false
	}
	for _, at := range user.AccessTokens {
		if !at.ExpireAt.After(now) {
			continue
		}
		if token.Token == at.Token {
			return at, true
		}
	}
	return nil, false
}

func (user *User) LookupAgentTokenByString(s string) (*AgentToken, bool) {
	token, err := ParseAgentTokenFromString(s)
	if err != nil {
		return nil, false
	}
	for _, at := range user.AgentTokens {
		if token.Token == at.Token {
			return at, true
		}
	}
	return nil, false
}
