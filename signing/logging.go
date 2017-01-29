package signing

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	. "github.com/PeerXu/jarvis3/signing/service"
	"github.com/PeerXu/jarvis3/user"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateUser(ctx context.Context, username string, password string, email string) (u *user.User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateUser",
			"username", username,
			"email", email,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.CreateUser(ctx, username, password, email)
}

func (s *loggingService) GetUserByID(ctx context.Context, id user.UserID) (u *user.User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetUserByID",
			"user_id", id.String(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.GetUserByID(ctx, id)
}

func (s *loggingService) Login(ctx context.Context, username string, password string) (u *user.User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Login",
			"username", username,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.Login(ctx, username, password)
}

func (s *loggingService) Logout(ctx context.Context, id user.UserID) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "Logout",
			"user_id", id.String(),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.Logout(ctx, id)
}

func (s *loggingService) CreateAgentToken(ctx context.Context, name string) (t *user.AgentToken, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateAgentToken",
			"name", name,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateAgentToken(ctx, name)
}

func (s *loggingService) DeleteAgentToken(ctx context.Context, name string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "DeleteAgentToken",
			"name", name,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteAgentToken(ctx, name)
}

func (s *loggingService) ValidateToken(ctx context.Context) (u *user.User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ValidateToken",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ValidateToken(ctx)
}
