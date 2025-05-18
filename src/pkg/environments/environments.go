package environments

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const (
	DEV = "development"
)

func LoadEnv() {
	appEnv := GetEnv("APP_ENV")

	if appEnv == "" {
		log.Error("Error loading APP_ENV")
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

func GetString(key string, defaultValue string) string {
	value := GetEnv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetInt(key string, defaultValue int) int {
	value := GetEnv(key)

	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return i
}

func GetBool(key string, defaultValue bool) bool {
	value := GetEnv(key)

	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return b
}

func IsDev() bool {
	appEnv := GetEnv("APP_ENV")
	return appEnv == DEV
}
