package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth/google"

	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/req"
)

type HandlerDeps struct {
	*configs.Config
	UserService *user.Service
}

type Handler struct {
	UserService *user.Service
	authService *Service
}

func NewHandler(app *fiber.App, deps HandlerDeps) {
	router := app.Group("/auth")
	// JWT Auth
	jwtService := NewService(ServiceDeps{
		UserService: deps.UserService,
	})
	handler := &Handler{
		UserService: deps.UserService,
		authService: jwtService,
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
		userInfo, err := handler.authService.Login(payload)

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
		data, err := handler.authService.register(payload)
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
	// Google Auth
	googleService := google.NewService(google.ServiceDeps{
		UserService: deps.UserService,
	})
	google.NewHandler(router, google.HandlerDeps{
		GoogleService: googleService,
	})
}
