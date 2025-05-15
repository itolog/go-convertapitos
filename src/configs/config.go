package configs

import (
	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"github.com/itolog/go-convertapitos/src/pkg/timeutils"
	"time"
)

type AuthConfig struct {
	JwtSecret           string `env:"JWT_SECRET"`
	CookieDomain        string `env:"COOKIE_DOMAIN" env-default:"localhost"`
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}

func LoadAuthConfig() (*AuthConfig, error) {
	accessTokenTTL := environments.GetEnv("JWT_ACCESS_TOKEN_TTL")
	refreshTokenTTL := environments.GetEnv("JWT_REFRESH_TOKEN_TTL")

	accessTokenDuration, err := timeutils.ParseDuration(accessTokenTTL)
	if err != nil {
		return nil, err
	}

	refreshTokenDuration, err := timeutils.ParseDuration(refreshTokenTTL)
	if err != nil {
		return nil, err
	}

	return &AuthConfig{
		JwtSecret:           environments.GetEnv("JWT_SECRET"),
		AccessTokenExpires:  accessTokenDuration,
		RefreshTokenExpires: refreshTokenDuration,
	}, nil
}

type DbConfig struct {
	Dsn string
}

type Config struct {
	Port   string `env:"PORT" env-default:"3000"`
	Prefix string `env:"PREFIX" env-default:"api"`
	Auth   *AuthConfig
	Db     DbConfig
}

func init() {
	environments.LoadEnv()
}

func NewConfig() *Config {
	authConfig, err := LoadAuthConfig()
	if err != nil {
		panic(err)
	}

	return &Config{
		Port:   environments.GetEnv("PORT"),
		Prefix: environments.GetEnv("PREFIX"),
		Auth:   authConfig,
		Db: DbConfig{
			Dsn: environments.GetEnv("DB_DSN"),
		},
	}
}
