package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/itolog/go-convertapitos/src/pkg/db"
)

type Service struct {
	UserRepository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		UserRepository: repository,
	}
}

func (service *Service) FindAll() ([]User, error) {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	return users, nil
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
