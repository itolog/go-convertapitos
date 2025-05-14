package configs

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

type AuthConfig struct {
	JwtSecret    string `env:"JWT_SECRET"`
	CookieDomain string `env:"COOKIE_DOMAIN" env-default:"localhost"`
}

type Config struct {
	Port   string `env:"PORT" env-default:"3000"`
	Prefix string `env:"PREFIX" env-default:"api"`
	Auth   AuthConfig
}

const (
	DEV  = "development"
	PROD = "production"
)

func init() {
	loadEnv()
}

func NewConfig() *Config {
	return &Config{
		Port:   GetConfigEnv("PORT"),
		Prefix: GetConfigEnv("PREFIX"),
		Auth: AuthConfig{
			JwtSecret:    GetConfigEnv("JWT_SECRET"),
			CookieDomain: GetConfigEnv("COOKIE_DOMAIN"),
		},
	}
}

func IsDev() bool {
	appEnv := GetConfigEnv("APP_ENV")
	return appEnv == DEV
}

func IsProd() bool {
	appEnv := GetConfigEnv("APP_ENV")
	return appEnv == PROD
}

func loadEnv() {
	appEnv := GetConfigEnv("APP_ENV")

	if appEnv == "" {
		panic("Error loading APP_ENV")
	}

	if IsDev() {
		err := godotenv.Load(".env.development")
		if err != nil {
			log.Error(err)
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Error(err)
		}
	}
}

func GetConfigEnv(key string) string {
	return os.Getenv(key)
}
