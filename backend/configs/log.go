package configs

import "github.com/itolog/go-convertapitos/backend/pkg/environments"

type LogConfig struct {
	Level  int
	Format string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  environments.GetInt("LOG_LEVEL", 0),
		Format: environments.GetString("LOG_FORMAT", "console"),
	}
}
