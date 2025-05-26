package user

import (
	"fmt"
	"github.com/itolog/go-convertapitos/backend/pkg/db"
	"gorm.io/gorm/clause"
)

type IRepository interface {
	Count() *int64
	FindAll(limit int, offset int, orderBy string, order string) ([]User, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
}

type Repository struct {
	Database *db.Db
}

const tableName = "users"

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) Count() *int64 {
	count := new(int64)

	repo.Database.DB.Table(tableName).Count(count)

	return count
}

func (repo *Repository) FindAll(limit int, offset int, orderBy string, order string) ([]User, error) {
	var users []User

	res := repo.Database.DB.
		Table(tableName).
		Omit("password").
		Order(fmt.Sprintf("%s %s", orderBy, order)).
		Limit(limit).
		Offset(offset).
		Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (repo *Repository) FindById(id string) (*User, error) {
	user := new(User)
	res := repo.Database.DB.First(user, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) FindByEmail(email string) (*User, error) {
	user := new(User)
	res := repo.Database.DB.First(user, "email = ?", email)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) Create(user *User) (*User, error) {
	res := repo.Database.DB.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) Update(user *User) (*User, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(user).Omit("password")

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) Delete(id string) error {
	res := repo.Database.DB.Where("id = ?", id).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
