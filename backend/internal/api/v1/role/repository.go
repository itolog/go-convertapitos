package role

import (
	"fmt"
	"github.com/itolog/go-convertapitos/backend/common/database"
	"github.com/itolog/go-convertapitos/backend/pkg/db"
	"gorm.io/gorm/clause"
)

const tableName = "roles"

type IRepository interface {
	database.ICrud[Role]
	GetForOptions(limit int, offset int, orderBy string, order string) ([]OptionsResponse, error)
}

type Repository struct {
	Database *db.Db
}

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

func (repo *Repository) FindAll(limit int, offset int, orderBy string, order string) ([]Role, error) {
	var roles []Role

	res := repo.Database.DB.
		Table(tableName).
		Order(fmt.Sprintf("%s %s", orderBy, order)).
		Limit(limit).
		Offset(offset).
		Find(&roles)

	if res.Error != nil {
		return nil, res.Error
	}

	return roles, nil
}

func (repo *Repository) GetForOptions(limit int, offset int, orderBy string, order string) ([]OptionsResponse, error) {
	var roles []OptionsResponse

	res := repo.Database.DB.
		Table(tableName).
		Omit("Permissions").
		Order(fmt.Sprintf("%s %s", orderBy, order)).
		Limit(limit).
		Offset(offset).
		Find(&roles)

	if res.Error != nil {
		return nil, res.Error
	}

	return roles, nil
}

func (repo *Repository) FindById(id string) (*Role, error) {
	user := new(Role)
	res := repo.Database.DB.First(user, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (repo *Repository) Create(role *Role) (*Role, error) {
	res := repo.Database.DB.Create(role)

	if res.Error != nil {
		return nil, res.Error
	}

	return role, nil
}

func (repo *Repository) Update(role *Role) (*Role, error) {
	selectFields := database.GetSelectFields[Role](role)

	res := repo.Database.DB.Clauses(clause.Returning{}).
		Select(selectFields).
		Updates(role)

	if res.Error != nil {
		return nil, res.Error
	}

	return role, nil
}

func (repo *Repository) Delete(id string) error {
	res := repo.Database.DB.Where("id = ?", id).Delete(&Role{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *Repository) BatchDelete(ids *[]string) error {
	res := repo.Database.DB.Where("ID IN (?)", *ids).Delete(&Role{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
