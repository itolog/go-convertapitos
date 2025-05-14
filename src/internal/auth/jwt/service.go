package jwt

import (
	"github.com/gofiber/fiber/v2"

	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/validation"
)

func (handler *HandlerJwtAuth) login(c *fiber.Ctx) error {
	validator := validation.NewValidator()
	payload := new(LoginRequest)

	err := c.BodyParser(payload)
	if err != nil {
		return err
	}
	validationErrors := validator.Validate(payload)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response[any]{
			Error: &api.ErrorResponse{
				Message: "Validation error",
				Code:    fiber.StatusBadRequest,
				Fields:  validationErrors,
			},
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
	return c.Status(fiber.StatusCreated).JSON(api.Response[LoginResponse]{
		Data: LoginResponse{
			AccessToken: "token",
		},
		Status: api.StatusSuccess,
	})
}
