package main

import (
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
	"github.com/itolog/go-convertapitos/backend/pkg/environments"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	environments.LoadEnv()
}

func main() {
	db, err := gorm.Open(postgres.Open(environments.GetEnv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}
}
