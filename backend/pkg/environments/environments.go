package environments

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

const (
	DEV = "development"
)

func LoadEnv() {
	if IsDev() {
		err := godotenv.Load(".env.development")
		if err != nil {
			log.Error().Msg(err.Error())
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Error().Msg(err.Error())
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
	appEnv := GetString("APP_ENV", "development")
	return appEnv == DEV
}
