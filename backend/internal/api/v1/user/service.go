package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/itolog/go-convertapitos/backend/pkg/db"
)

type IUserService interface {
	FindAll(limit int, offset int, orderBy string, desc bool) (*FindAllResponse, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user User) (*User, error)
	Update(id string, payload *UpdateRequest) (*User, error)
	Delete(id string) error
	BatchDelete(ids *[]string) error
	CreateOrUpdateAccount(userID uuid.UUID, account Account) error
	FindByProviderAccount(provider, providerID string) (*User, error)
}

type Service struct {
	UserRepository IRepository
}

func NewService(repository *Repository) *Service {
	return &Service{
		UserRepository: repository,
	}
}

func (service *Service) FindAll(limit int, offset int, orderBy string, desc bool) (*FindAllResponse, error) {
	count := service.UserRepository.Count()

	order := "asc"
	if desc {
		order = "desc"
	}
	users, err := service.UserRepository.FindAll(limit, offset, orderBy, order)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	return &FindAllResponse{
		Users: users,
		Count: count,
	}, nil
}

func (service *Service) FindById(id string) (*User, error) {
	user, err := service.UserRepository.FindById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return user, nil
}

func (service *Service) FindByEmail(email string) (*User, error) {
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return user, nil
}

func (service *Service) Create(user User) (*User, error) {
	created, err := service.UserRepository.Create(&user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	return created, nil
}

func (service *Service) Update(id string, payload *UpdateRequest) (*User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	updated, err := service.UserRepository.Update(&User{
		Model:         db.Model{ID: uid},
		Name:          payload.Name,
		Email:         payload.Email,
		VerifiedEmail: payload.VerifiedEmail,
		Picture:       payload.Picture,
		Password:      payload.Password,
		RoleID:        payload.RoleID,
		AuthMethod:    payload.AuthMethod,
	})

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	return updated, nil
}

func (service *Service) Delete(id string) error {
	_, err := service.UserRepository.FindById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	err = service.UserRepository.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (service *Service) BatchDelete(ids *[]string) error {
	err := service.UserRepository.BatchDelete(ids)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (service *Service) CreateOrUpdateAccount(userID uuid.UUID, account Account) error {
	account.UserID = userID
	err := service.UserRepository.CreateOrUpdateAccount(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return nil
}

func (service *Service) FindByProviderAccount(provider, providerID string) (*User, error) {
	user, err := service.UserRepository.FindByProviderAccount(provider, providerID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return user, nil
}
