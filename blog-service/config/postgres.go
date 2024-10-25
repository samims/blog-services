package config

import (
	"blog-service/constants"
	"fmt"

	"github.com/spf13/viper"
)

type PostgresConfig interface {
	Host() string
	Port() string
	User() string
	Password() string
	Database() string
	SSLMode() string

	ConnectionURL() string
}
type postgresConfig struct {
	env *viper.Viper
}

func (cfg *postgresConfig) Host() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.DatabaseHost)
}

func (cfg *postgresConfig) Port() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.DatabasePort)
}

func (cfg *postgresConfig) User() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.DatabaseUser)
}

func (cfg *postgresConfig) Password() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.DatabasePasswd)
}

func (cfg *postgresConfig) Database() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.DatabaseName)
}

func (cfg *postgresConfig) SSLMode() string {
	cfg.env.AutomaticEnv()
	// default "disabled"
	var sslMode string
	sslMode = cfg.env.GetString(constants.DatabaseSSL)
	if sslMode == "" {
		sslMode = constants.DatabaseSSL
	}
	return sslMode
}

func (cfg *postgresConfig) ConnectionURL() string {
	cfg.env.AutomaticEnv()
	connectionURL := fmt.Sprintf(
		`%s:%s@%s:%s/%s?sslmode=%s`,
		cfg.User(),
		cfg.Password(),
		cfg.Host(),
		cfg.Port(),
		cfg.Database(),
		cfg.SSLMode(),
	)
	return connectionURL
}

func NewPostgresConfig(env *viper.Viper) PostgresConfig {
	return &postgresConfig{
		env: env,
	}
}
