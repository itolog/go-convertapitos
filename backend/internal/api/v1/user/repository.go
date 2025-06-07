package user

import (
	"errors"
	"fmt"
	"github.com/itolog/go-convertapitos/backend/pkg/db"
	"gorm.io/gorm"
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
	BatchDelete(ids *[]string) error
	CreateOrUpdateAccount(account *Account) error
	FindByProviderAccount(provider, providerID string) (*User, error)
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
		Preload(clause.Associations).
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
	selectFields := getSelectFields(user)

	res := repo.Database.DB.Clauses(clause.Returning{}).
		Select(selectFields).
		Updates(user).Omit("password")

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

func (repo *Repository) BatchDelete(ids *[]string) error {
	res := repo.Database.DB.Where("ID IN (?)", *ids).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// CreateOrUpdateAccount Creates or updates an account
func (repo *Repository) CreateOrUpdateAccount(account *Account) error {
	existingAccount := new(Account)

	// searching for an existing account
	result := repo.Database.DB.Where("provider = ? AND provider_id = ?", account.Provider, account.ProviderID).First(&existingAccount)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// create a new account
		if err := repo.Database.DB.Create(account).Error; err != nil {
			return err
		}
	} else if result.Error != nil {
		return result.Error
	} else {
		// update the existing
		existingAccount.AccessToken = account.AccessToken
		existingAccount.RefreshToken = account.RefreshToken
		existingAccount.ExpiresAt = account.ExpiresAt
		if err := repo.Database.DB.Save(&existingAccount).Error; err != nil {
			return err
		}
	}
	return nil
}

// FindByProviderAccount Finds the user by provider account
func (repo *Repository) FindByProviderAccount(provider, providerID string) (*User, error) {
	var account Account
	if err := repo.Database.DB.Where("provider = ? AND provider_id = ?", provider, providerID).First(&account).Error; err != nil {
		return nil, err
	}

	var user User
	if err := repo.Database.DB.Where("id = ?", account.UserID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
