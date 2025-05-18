package db

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/itolog/go-convertapitos/src/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	return &Db{
		db,
	}
}
