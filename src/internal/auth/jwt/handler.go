package jwt

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/req"
)

type HandlerDeps struct {
	JwtService *Service
}

type Handler struct {
	JwtService *Service
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := Handler{

		JwtService: deps.JwtService,
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
		userInfo, err := handler.JwtService.login(payload)

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
		data, err := handler.JwtService.register(payload)
		if err != nil {
			var fiberErr *fiber.Error
			ok := errors.As(err, &fiberErr)
			statusCode := fiber.StatusInternalServerError

			if ok {
				statusCode = fiberErr.Code
			}

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
