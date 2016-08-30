package signing

import (
	"errors"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	"github.com/PeerXu/jarvis3/user"
)

type Service interface {
	CreateUser(ctx context.Context, username string, password string, email string) (*user.User, error)
	DeleteUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (*user.User, error)
	Login(ctx context.Context, username string, password string) (*user.AccessToken, error)
	Logout(ctx context.Context, username string) error
	CreateAgentToken(ctx context.Context, name string) (*user.AgentToken, error)
	DeleteAgentToken(ctx context.Context, name string) error
	ValidateAccessToken(ctx context.Context) (*user.User, error)
}

type service struct {
	logger         log.Logger
	userRepository user.Repository
}

func NewService(logger log.Logger, userRepository user.Repository) Service {
	return &service{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (s *service) CreateUser(ctx context.Context, username string, password string, email string) (*user.User, error) {
	u := user.NewUser(username, password, email)
	_, err := s.userRepository.CreateUser(u)
	if err != nil {
		return nil, newSignError(errorServerError, "failed to create user", err)
	}
	return u, nil
}

func (s *service) DeleteUser(ctx context.Context, username string) error {
	err := s.userRepository.DeleteUser(username)
	if err != nil {
		return newSignError(errorServerError, "failed to delete user", err)
	}
	return nil
}

func (s *service) GetUser(ctx context.Context, username string) (*user.User, error) {
	u, err := s.userRepository.GetUser(username)
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			return nil, newSignError(errorNotFound, "user not found", err)
		default:
			return nil, newSignError(errorServerError, "failed to get user", err)
		}
	}
	return u, nil
}

func (s *service) Login(ctx context.Context, username string, password string) (*user.AccessToken, error) {
	u, err := s.userRepository.GetUser(username)
	if err != nil {
		return nil, newSignError(errorAccessDenied, "access denied", user.ErrInvalidUsernameOrPassword)
	}

	if !u.ValidatePassword(password) {
		return nil, newSignError(errorAccessDenied, "access denied", user.ErrInvalidUsernameOrPassword)
	}

	token := user.NewAccessToken()
	err = s.userRepository.CreateAccessToken(u, token)
	if err != nil {
		return nil, newSignError(errorServerError, "failed to create access token", err)
	}

	return token, nil
}

func (s *service) Logout(ctx context.Context, username string) error {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	if u.Username != username {
		return newSignError(errorAccessDenied, "access denied", errors.New("access token don't match user"))
	}

	t := jctx.AccessToken()
	at := &user.AccessToken{Token: t.Token}
	err := s.userRepository.DeleteAccessTokens(u, []*user.AccessToken{at})
	if err != nil {
		return newSignError(errorServerError, "failed to delete access token", err)
	}
	return nil
}

func (s *service) CreateAgentToken(ctx context.Context, name string) (*user.AgentToken, error) {
	u := ctx.Value("JarvisContext").(jcontext.Context).User()
	t := user.NewAgentToken(name)

	err := s.userRepository.CreateAgentToken(u, t)
	if err != nil {
		return nil, newSignError(errorServerError, "failed to create agent token", err)
	}

	return t, nil
}

func (s *service) DeleteAgentToken(ctx context.Context, name string) error {
	u := ctx.Value("JarvisContext").(jcontext.Context).User()

	t, err := s.userRepository.LookupAgentTokenByName(u, name)
	if err != nil {
		return newSignError(errorServerError, "failed to lookup agent token", err)
	}
	err = s.userRepository.DeleteAgentTokens(u, []*user.AgentToken{t})
	if err != nil {
		return newSignError(errorServerError, "failed to delete agent token", err)
	}
	return nil
}

func (s *service) ValidateAccessToken(ctx context.Context) (*user.User, error) {
	tokenStr := ctx.Value("Authorization").(string)
	token := &user.AccessToken{Token: tokenStr}
	u, err := s.userRepository.LookupUserByAccessToken(token)
	if err != nil {
		return nil, newSignError(errorAccessDenied, "access denied", err)
	}
	var ts []*user.AccessToken
	for _, t := range u.AccessTokens {
		if t.Token == tokenStr {
			ts = append(ts, t)
			break
		}
	}
	u.AccessTokens = ts
	return u, nil
}
