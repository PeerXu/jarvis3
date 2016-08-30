package user

import "time"

type User struct {
	Username string
	Password EncryptedPassword
	Email    Email

	AccessTokens []*AccessToken
	AgentTokens  []*AgentToken
}

func NewUser(username, password, email string) *User {
	return &User{
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

func (user *User) ValidateAccessToken(token *AccessToken) bool {
	now := time.Now()
	for _, at := range user.AccessTokens {
		if !at.ExpireAt.After(now) {
			continue
		}
		if token.Token == at.Token {
			return true
		}
	}
	return false
}
