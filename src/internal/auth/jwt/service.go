package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/req"

	"github.com/itolog/go-convertapitos/src/pkg/api"
)

type ServiceDeps struct {
	*configs.Config
	UserRepository *user.Repository
}
type Service struct {
	*configs.Config
	UserRepository *user.Repository
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserRepository: deps.UserRepository,
		Config:         deps.Config,
	}
}

func (service *Service) login(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[LoginRequest](c)
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

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: LoginResponse{
			AccessToken: "token",
		},
		Status: api.StatusSuccess,
	})
}

func (service *Service) register(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[RegisterRequest](c)
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

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: RegisterResponse{
			AccessToken: "reg",
		},
		Status: api.StatusSuccess,
	})
}
