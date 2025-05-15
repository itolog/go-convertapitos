package environments

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

const (
	DEV = "development"
)

func LoadEnv() {
	appEnv := GetEnv("APP_ENV")

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

func GetEnv(key string) string {
	return os.Getenv(key)
}

func IsDev() bool {
	appEnv := GetEnv("APP_ENV")
	return appEnv == DEV
}
