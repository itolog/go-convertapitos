package user

import (
	"github.com/itolog/go-convertapitos/src/pkg/db"
	"gorm.io/gorm/clause"
)

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) findAll() ([]User, error) {
	var users []User
	res := repo.Database.DB.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (repo *Repository) findById(id string) (*User, error) {
	user := new(User)
	res := repo.Database.DB.First(user, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) create(user *User) (*User, error) {
	res := repo.Database.DB.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) update(user *User) (*User, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(user).Omit("password")

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) delete(id string) error {
	res := repo.Database.DB.Where("id = ?", id).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
