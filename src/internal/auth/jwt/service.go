package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/api"
)

func (handler *HandlerJwtAuth) login(c *fiber.Ctx) error {

	payload := new(LoginRequest)

	err := c.BodyParser(payload)
	if err != nil {
		return err
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
