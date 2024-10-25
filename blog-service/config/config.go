package config

import "github.com/spf13/viper"

// Configuration is the blueprint of configuration
type Configuration interface {
	AppConfig() AppConfig
	PostgresConfig() PostgresConfig
}

// configuration holds the required config instance
type configuration struct {
	appConfig      AppConfig
	postgresConfig PostgresConfig
}

func (c *configuration) AppConfig() AppConfig {
	return c.appConfig
}

func (c *configuration) PostgresConfig() PostgresConfig {
	return c.postgresConfig
}

func Init(v *viper.Viper) Configuration {
	return &configuration{
		appConfig:      NewAppConfig(v),
		postgresConfig: NewPostgresConfig(v),
	}
}
