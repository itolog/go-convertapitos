package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string `env:"PORT" env-default:"3000"`
	Prefix       string `env:"PREFIX" env-default:"api"`
	CookieDomain string `env:"COOKIE_DOMAIN" env-default:"localhost"`
}

const (
	DEV  = "development"
	PROD = "production"
)

func NewConfig() *Config {
	appEnv := GetConfigEnv("APP_ENV")

	if appEnv == "" {
		panic("Error loading APP_ENV")
	}

	if IsDev() {
		err := godotenv.Load(".env.development")
		if err != nil {
			fmt.Println("Error loading .env.development file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
		}

	}

	port := os.Getenv("PORT")
	prefix := os.Getenv("PREFIX")
	cookieDomain := os.Getenv("COOKIE_DOMAIN")

	return &Config{
		Port:         port,
		Prefix:       prefix,
		CookieDomain: cookieDomain,
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

func GetConfigEnv(key string) string {
	return os.Getenv(key)
}
