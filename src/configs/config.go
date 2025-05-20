package configs

import (
	"github.com/itolog/go-convertapitos/src/pkg/environments"
)

type DbConfig struct {
	Dsn string
}

type Config struct {
	Port          string `env:"PORT" env-default:"3000"`
	EnablePrefork bool   `env:"ENABLE_PREFORK" env-default:"false"`
	Db            DbConfig
}

func init() {
	environments.LoadEnv()
}

func NewConfig() *Config {
	return &Config{
		Port:          environments.GetString("PORT", "3000"),
		EnablePrefork: environments.GetBool("ENABLE_PREFORK", false),
		Db: DbConfig{
			Dsn: environments.GetEnv("DB_DSN"),
		},
	}
}
