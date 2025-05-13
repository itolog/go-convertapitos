package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string `env:"PORT" env-default:"3000"`
	Prefix string `env:"PREFIX" env-default:"api"`
}

const (
	DEV  = "development"
	PROD = "production"
)

func NewConfig() *Config {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		panic("Error loading APP_ENV")
	}

	if appEnv != PROD {
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

	return &Config{
		Port:   port,
		Prefix: prefix,
	}
}

func GetConfigEnv(key string) string {
	return os.Getenv(key)
}
