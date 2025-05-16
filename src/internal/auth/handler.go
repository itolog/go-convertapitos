package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/req"
)

type HandlerDeps struct {
	AuthService *Service
}

type Handler struct {
	AuthService *Service
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := &Handler{
		AuthService: deps.AuthService,
	}

	router.Post("/login", func(c *fiber.Ctx) error {
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
		userInfo, err := handler.AuthService.Login(payload)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return c.Status(fiber.StatusOK).JSON(api.Response{
			Data:   userInfo,
			Status: api.StatusSuccess,
		})
	})
	router.Post("/register", func(c *fiber.Ctx) error {
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
		data, err := handler.AuthService.register(payload)
		if err != nil {
			statusCode := api.GetErrorCode(err)

			return c.Status(statusCode).JSON(api.Response{
				Error: &api.ErrorResponse{
					Message: err.Error(),
				},
				Status: api.StatusError,
			})
		}

		return c.Status(fiber.StatusOK).JSON(api.Response{
			Data:   data,
			Status: api.StatusSuccess,
		})
	})
}
