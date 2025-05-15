package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/db"
	"github.com/itolog/go-convertapitos/src/pkg/req"
)

type Service struct {
	UserRepository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		UserRepository: repository,
	}
}

func (service *Service) findAll(c *fiber.Ctx) error {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data:   users,
		Status: api.StatusSuccess,
	})
}

func (service *Service) findById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := service.UserRepository.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data:   user,
		Status: api.StatusSuccess,
	})
}

func (service *Service) findByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data:   user,
		Status: api.StatusSuccess,
	})
}

func (service *Service) create(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[CreateRequest](c)
	if err != nil {
		return err
	}
	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	user := User{
		Name:          payload.Name,
		Email:         payload.Email,
		VerifiedEmail: payload.VerifiedEmail,
		Picture:       payload.Picture,
		Password:      payload.Password,
	}

	created, err := service.UserRepository.Create(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data:   created,
		Status: api.StatusSuccess,
	})
}

func (service *Service) update(c *fiber.Ctx) error {
	id := c.Params("id")

	payload, err := req.DecodeBody[UpdateRequest](c)
	if err != nil {
		return err
	}

	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
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
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data:   updated,
		Status: api.StatusSuccess,
	})
}

func (service *Service) delete(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := service.UserRepository.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: "User not found",
			},
			Status: api.StatusError,
		})
	}

	err = service.UserRepository.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data:   fmt.Sprintf("User with id %s deleted", id),
		Status: api.StatusSuccess,
	})
}
