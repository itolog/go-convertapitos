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

	router.Post("/login", func(ctx *fiber.Ctx) error {
		payload, err := req.DecodeBody[LoginRequest](ctx)
		if err != nil {
			return err
		}

		validateError, valid := req.ValidateBody(payload)
		if !valid {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error:  validateError,
				Status: api.StatusError,
			})
		}

		userInfo, err := handler.AuthService.Login(ctx, payload)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   userInfo,
			Status: api.StatusSuccess,
		})
	})
	router.Post("/register", func(ctx *fiber.Ctx) error {
		payload, err := req.DecodeBody[RegisterRequest](ctx)
		if err != nil {
			return err
		}

		validateError, valid := req.ValidateBody(payload)
		if !valid {
			return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
				Error:  validateError,
				Status: api.StatusError,
			})
		}

		data, err := handler.AuthService.register(ctx, payload)
		if err != nil {
			statusCode := api.GetErrorCode(err)
			return fiber.NewError(statusCode, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   data,
			Status: api.StatusSuccess,
		})
	})
	router.Get("/refresh-token", func(ctx *fiber.Ctx) error {
		refreshToken := ctx.Cookies("refreshToken")
		if refreshToken == "" {
			return fiber.NewError(fiber.StatusUnauthorized, api.ErrUnauthorized)
		}

		user, err := handler.AuthService.RefreshToken(ctx, refreshToken)
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).JSON(api.Response{
			Data:   user,
			Status: api.StatusSuccess,
		})
	})
}
