package config

import (
	"fmt"

	"blog-service/constants"

	_ "github.com/lib/pq"
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
	ConnectionURLWithScheme() string
}
type postgresConfig struct {
	env *viper.Viper
}

func (cfg *postgresConfig) Host() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.PostgresHost)
}

func (cfg *postgresConfig) Port() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.PostgresPort)
}

func (cfg *postgresConfig) User() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.PostgresUser)
}

func (cfg *postgresConfig) Password() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.PostgresPasswd)
}

func (cfg *postgresConfig) Database() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString(constants.PostgresDBName)
}

func (cfg *postgresConfig) SSLMode() string {
	cfg.env.AutomaticEnv()
	// default "disabled"
	var sslMode string
	sslMode = cfg.env.GetString(constants.DatabaseSSLMode)
	if sslMode == "" {
		sslMode = constants.DatabaseDefaultSSL
	}
	return sslMode
}

func (cfg *postgresConfig) ConnectionURL() string {
	cfg.env.AutomaticEnv()

	pgInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		cfg.Host(),
		cfg.Port(),
		cfg.User(),
		cfg.Password(),
		cfg.Database())

	return pgInfo
}

// ConnectionURLWithScheme returns the connection string in URL format with postgres:// scheme
func (cfg *postgresConfig) ConnectionURLWithScheme() string {
	cfg.env.AutomaticEnv()
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User(),
		cfg.Password(),
		cfg.Host(),
		cfg.Port(),
		cfg.Database())
}

func NewPostgresConfig(env *viper.Viper) PostgresConfig {
	return &postgresConfig{
		env: env,
	}
}
