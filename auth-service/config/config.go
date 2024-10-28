package config

import (
	"auth-service/constants"

	"github.com/spf13/viper"
)

// Configuration is the blueprint of configuration
type Configuration interface {
	AppConfig() AppConfig
	Load(env *viper.Viper)
}

type configuration struct {
	appConfig AppConfig
}

func (c *configuration) AppConfig() AppConfig {
	return c.appConfig
}

func (c *configuration) Load(environ *viper.Viper) {
	environ.AutomaticEnv()
}

func NewConfiguration(appConf AppConfig) Configuration {
	return &configuration{appConfig: appConf}
}

type AppConfig interface {
	BuildEnv() string
	SecretKey() string
	Port() string
}

type appConfig struct {
	env *viper.Viper
}

// NewAppConfig gives a new instance of appConfig struct
func NewAppConfig(env *viper.Viper) AppConfig {
	return &appConfig{
		env: env,
	}
}

// BuildEnv returns environment type
func (ac *appConfig) BuildEnv() string {
	ac.env.AutomaticEnv()
	return ac.env.GetString(constants.BuildEnv)
}

func (ac *appConfig) SecretKey() string {
	ac.env.AutomaticEnv()
	return ac.env.GetString(constants.SecretKey)
}

func (ac *appConfig) Port() string {
	ac.env.AutomaticEnv()
	return ac.env.GetString(constants.AppPort)
}
