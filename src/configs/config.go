package configs

import (
	"github.com/itolog/go-convertapitos/src/pkg/environments"
)

type DbConfig struct {
	Dsn string
}

type Config struct {
	Port string `env:"PORT" env-default:"3000"`
	Db   DbConfig
}

func init() {
	environments.LoadEnv()
}

func NewConfig() *Config {
	return &Config{
		Port: environments.GetString("PORT", "3000"),
		Db: DbConfig{
			Dsn: environments.GetEnv("DB_DSN"),
		},
	}
}
