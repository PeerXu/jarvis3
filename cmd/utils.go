package cmd

import (
	"github.com/go-kit/kit/log"
	"github.com/spf13/viper"

	"github.com/PeerXu/jarvis3/signing"
	signing_client "github.com/PeerXu/jarvis3/signing/client"
	"github.com/PeerXu/jarvis3/utils"
)

type Configure struct {
	Host        string
	Username    string
	AccessToken string
}

func LoadConfigure() Configure {
	return Configure{
		Host:        viper.GetString("host"),
		Username:    viper.GetString("username"),
		AccessToken: viper.GetString("access_token"),
	}
}

func LoadEnvironmentFromConfigure(cfg Configure) utils.Environment {
	env := utils.NewEnvironment()

	env.Set("Username", cfg.Username)
	env.Set("AccessToken", cfg.AccessToken)

	return env
}

func SimpleNewSigningClient(logger log.Logger) (signing.Service, error) {
	cfg := LoadConfigure()
	return signing_client.New(cfg.Host, utils.NewEnvironment(), logger)
}

func SimpleNewSigningClientWithEnvironment(logger log.Logger) (signing.Service, error) {
	cfg := LoadConfigure()
	return signing_client.New(cfg.Host, LoadEnvironmentFromConfigure(cfg), logger)
}
