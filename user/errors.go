package user

import "errors"

var (
	ErrUnknown                   = errors.New("user unknown error")
	ErrUserNotFound              = errors.New("user not found")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrAccessTokenNotFound       = errors.New("access token not found")
	ErrAgentTokenNotFound        = errors.New("agent token not found")
	ErrBadTokenFormat            = errors.New("bad format token")
)
