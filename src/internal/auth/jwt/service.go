package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/req"

	"github.com/itolog/go-convertapitos/src/pkg/api"
)

func (handler *HandlerJwtAuth) login(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[LoginRequest](c)
	if err != nil {
		return err
	}
	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response[any]{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.Response[LoginResponse]{
		Data: LoginResponse{
			AccessToken: "token",
		},
		Status: api.StatusSuccess,
	})
}

func (handler *HandlerJwtAuth) register(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[RegisterRequest](c)
	if err != nil {
		return err
	}
	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response[any]{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response[RegisterResponse]{
		Data: RegisterResponse{
			AccessToken: "reg",
		},
		Status: api.StatusSuccess,
	})
}
