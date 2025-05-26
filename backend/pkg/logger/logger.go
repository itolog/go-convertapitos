package logger

import (
	"github.com/itolog/go-convertapitos/backend/configs"
	"github.com/rs/zerolog"
	"os"
)

func NewLogger(config *configs.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
	var logger zerolog.Logger

	if config.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	return &logger
}
