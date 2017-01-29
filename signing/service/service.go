package service

import (
	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/user"
)

type Service interface {
	CreateUser(ctx context.Context, username string, password string, email string) (*user.User, error)
	DeleteUserByID(ctx context.Context, id user.UserID) error
	GetUserByID(ctx context.Context, id user.UserID) (*user.User, error)
	Login(ctx context.Context, username string, password string) (*user.User, error)
	Logout(ctx context.Context, id user.UserID) error
	CreateAgentToken(ctx context.Context, name string) (*user.AgentToken, error)
	DeleteAgentToken(ctx context.Context, name string) error
	ValidateToken(ctx context.Context) (*user.User, error)
}
