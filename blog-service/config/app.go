package config

import (
	"blog-service/constants"

	"github.com/spf13/viper"
)

type AppConfig interface {
	GetBuildEnv() string
	GetSecretKey() string
}

// appConfig for app
type appConfig struct {
	env *viper.Viper
}

// GetBuildEnv returns environment type
func (ac *appConfig) GetBuildEnv() string {
	ac.env.AutomaticEnv()
	return ac.env.GetString(constants.SecretKey)
}

func (ac *appConfig) GetSecretKey() string {
	ac.env.AutomaticEnv()
	return ac.env.GetString(constants.SecretKey)
}

func NewAppConfig(env *viper.Viper) AppConfig {
	return &appConfig{env: env}
}
