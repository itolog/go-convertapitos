package configs

import (
	"github.com/itolog/go-convertapitos/src/pkg/environments"
)

type AuthConfig struct {
	JwtSecret    string `env:"JWT_SECRET"`
	CookieDomain string `env:"COOKIE_DOMAIN" env-default:"localhost"`
}

type DbConfig struct {
	Dsn string
}

type Config struct {
	Port   string `env:"PORT" env-default:"3000"`
	Prefix string `env:"PREFIX" env-default:"api"`
	Auth   AuthConfig
	Db     DbConfig
}

func init() {
	environments.LoadEnv()
}

func NewConfig() *Config {
	return &Config{
		Port:   environments.GetEnv("PORT"),
		Prefix: environments.GetEnv("PREFIX"),
		Auth: AuthConfig{
			JwtSecret:    environments.GetEnv("JWT_SECRET"),
			CookieDomain: environments.GetEnv("COOKIE_DOMAIN"),
		},
		Db: DbConfig{
			Dsn: environments.GetEnv("DB_DSN"),
		},
	}
}
