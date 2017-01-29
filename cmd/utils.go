package cmd

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/spf13/viper"

	computing_client "github.com/PeerXu/jarvis3/computing/client"
	computing_service "github.com/PeerXu/jarvis3/computing/service"
	signing_client "github.com/PeerXu/jarvis3/signing/client"
	signing_service "github.com/PeerXu/jarvis3/signing/service"
	"github.com/PeerXu/jarvis3/utils"
)

func NewClientLogger() log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	return logger
}

type Configure struct {
	Host        string
	UserID      string
	Username    string
	AccessToken string
}

type SigningConfigure struct {
	Host        string
	UserID      string
	AccessToken string
}

type ComputingConfigure struct {
	Host string
}

func LoadConfigure() Configure {
	return Configure{
		Host:        viper.GetString("host"),
		UserID:      viper.GetString("user_id"),
		Username:    viper.GetString("username"),
		AccessToken: viper.GetString("access_token"),
	}
}

func LoadEnvironmentFromConfigure(cfg Configure) utils.Environment {
	env := utils.NewEnvironment()

	env.Set("Username", cfg.Username)
	env.Set("AccessToken", cfg.AccessToken)
	env.Set("UserID", cfg.UserID)

	return env
}

func NewSigningClientForClientWithoutEnvironment(logger log.Logger) (signing_service.Service, error) {
	cfg := LoadConfigure()
	return signing_client.NewForClient(cfg.Host, utils.NewEnvironment(), logger)
}

func NewSigningClientForClient(logger log.Logger) (signing_service.Service, error) {
	cfg := LoadConfigure()
	return signing_client.NewForClient(cfg.Host, LoadEnvironmentFromConfigure(cfg), logger)
}

func NewSigningClientForService(logger log.Logger) (signing_service.Service, error) {
	instance := viper.GetString("development.service.signing.host")
	return signing_client.NewForService(instance, logger)
}

func NewComputingClientForClient(logger log.Logger) (computing_service.Service, error) {
	cfg := LoadConfigure()
	return computing_client.NewForClient(cfg.Host, LoadEnvironmentFromConfigure(cfg), logger)
}
