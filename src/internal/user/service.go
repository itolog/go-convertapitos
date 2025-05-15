package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/db"
	"github.com/itolog/go-convertapitos/src/pkg/req"
)

func (handler *Handler) findAll(c *fiber.Ctx) error {
	users, err := handler.repository.findAll()
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

func (handler *Handler) findById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := handler.repository.findById(id)
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

func (handler *Handler) create(c *fiber.Ctx) error {
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

	created, err := handler.repository.create(&user)
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

func (handler *Handler) update(c *fiber.Ctx) error {
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

	updated, err := handler.repository.update(&User{
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

func (handler *Handler) delete(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := handler.repository.findById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: "User not found",
			},
			Status: api.StatusError,
		})
	}

	err = handler.repository.delete(id)
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
