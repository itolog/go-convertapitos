package user

import "github.com/itolog/go-convertapitos/src/pkg/db"

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) create(user *User) {

}
