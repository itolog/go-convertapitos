package user

import (
	"github.com/itolog/go-convertapitos/src/pkg/db"
	"gorm.io/gorm/clause"
)

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

func (repo *Repository) FindAll(limit, offset int) ([]User, error) {
	var users []User
	res := repo.Database.DB.
		Table(tableName).
		Omit("password").
		Order("updated_at desc").
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
