package signing

import (
	"errors"
	"strings"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	jerrors "github.com/PeerXu/jarvis3/errors"
	. "github.com/PeerXu/jarvis3/signing/error"
	. "github.com/PeerXu/jarvis3/signing/service"
	"github.com/PeerXu/jarvis3/user"
)

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
		return nil, NewSignError(jerrors.ErrorServerError, "failed to create user", err)
	}
	return u, nil
}

func (s *service) DeleteUserByID(ctx context.Context, id user.UserID) error {
	err := s.userRepository.DeleteUserByID(id)
	if err != nil {
		err = NewSignError(jerrors.ErrorServerError, "failed to delete user", err)
		return err
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id user.UserID) (*user.User, error) {
	u, err := s.userRepository.GetUserByID(id)
	if err != nil {
		switch err {
		case user.ErrUserNotFound:
			err = NewSignError(jerrors.ErrorNotFound, "user not found", err)
		default:
			err = NewSignError(jerrors.ErrorServerError, "failed to get user", err)
		}
		return nil, err
	}
	return u, nil
}

func (s *service) Login(ctx context.Context, username string, password string) (*user.User, error) {
	u, err := s.userRepository.LookupUserByUsername(username)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", user.ErrInvalidUsernameOrPassword)
	}

	if !u.ValidatePassword(password) {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", user.ErrInvalidUsernameOrPassword)
	}

	token := user.NewAccessToken()
	err = s.userRepository.CreateAccessToken(u.ID, token)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorServerError, "failed to create access token", err)
	}

	return &user.User{
		ID:           u.ID,
		Username:     u.Username,
		Email:        u.Email,
		AccessTokens: []*user.AccessToken{token},
	}, nil
}

func (s *service) Logout(ctx context.Context, id user.UserID) error {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	if u.ID != id {
		return NewSignError(jerrors.ErrorAccessDenied, "access denied", errors.New("access token don't match user"))
	}

	t := jctx.AccessToken()
	at := &user.AccessToken{Token: t.Token}
	err := s.userRepository.DeleteAccessTokens(u.ID, []*user.AccessToken{at})
	if err != nil {
		return NewSignError(jerrors.ErrorServerError, "failed to delete access token", err)
	}
	return nil
}

func (s *service) CreateAgentToken(ctx context.Context, name string) (*user.AgentToken, error) {
	u := ctx.Value("JarvisContext").(jcontext.Context).User()
	t := user.NewAgentToken(name)

	err := s.userRepository.CreateAgentToken(u.ID, t)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorServerError, "failed to create agent token", err)
	}

	return t, nil
}

func (s *service) DeleteAgentToken(ctx context.Context, name string) error {
	u := ctx.Value("JarvisContext").(jcontext.Context).User()

	t, err := s.userRepository.LookupAgentTokenByName(u.ID, name)
	if err != nil {
		return NewSignError(jerrors.ErrorServerError, "failed to lookup agent token", err)
	}
	err = s.userRepository.DeleteAgentTokens(u.ID, []*user.AgentToken{t})
	if err != nil {
		return NewSignError(jerrors.ErrorServerError, "failed to delete agent token", err)
	}
	return nil
}

func (s *service) validateAccessToken(ctx context.Context) (*user.User, error) {
	tokenStr := ctx.Value("Authorization").(string)
	token, err := user.ParseAccessTokenFromString(tokenStr)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", err)
	}

	u, err := s.userRepository.LookupUserByAccessToken(token)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", err)
	}

	var ts []*user.AccessToken
	for _, t := range u.AccessTokens {
		if t.Token == token.Token {
			ts = append(ts, t)
			break
		}
	}
	u.AccessTokens = ts
	return u, nil
}

func (s *service) validateAgentToken(ctx context.Context) (*user.User, error) {
	tokenStr := ctx.Value("Authorization").(string)
	token, err := user.ParseAgentTokenFromString(tokenStr)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", err)
	}

	u, err := s.userRepository.LookupUserByAgentToken(token)
	if err != nil {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", err)
	}

	var ts []*user.AgentToken
	for _, t := range u.AgentTokens {
		if t.Token == token.Token {
			ts = append(ts, t)
			break
		}
	}
	u.AgentTokens = ts
	return u, nil
}

func (s *service) ValidateToken(ctx context.Context) (*user.User, error) {
	tokenStr := ctx.Value("Authorization").(string)

	if len(tokenStr) < 3 {
		return nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", user.ErrBadTokenFormat)
	}

	var u *user.User
	var err error
	switch strings.ToLower(tokenStr[:3]) {
	case "act":
		u, err = s.validateAccessToken(ctx)
	case "agt":
		u, err = s.validateAgentToken(ctx)
	default:
		u, err = nil, NewSignError(jerrors.ErrorAccessDenied, "access denied", user.ErrBadTokenFormat)
	}

	return u, err
}
