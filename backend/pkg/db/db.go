package db

import (
	"github.com/itolog/go-convertapitos/backend/configs"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Error().Msg("Error while connecting to database")
	}
	return &Db{
		db,
	}
}
