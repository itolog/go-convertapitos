package google

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/api"
)

type HandlerDeps struct {
	GoogleService *Service
}

type Handler struct {
	GoogleService *Service
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := Handler{
		GoogleService: deps.GoogleService,
	}

	router.Get("/google", func(ctx *fiber.Ctx) error {
		from := ctx.Query("from", "/")
		path := configs.ConfigGoogle()
		url := path.AuthCodeURL(from)

		return ctx.Redirect(url)
	})

	router.Get("/google/callback", func(ctx *fiber.Ctx) error {
		code := ctx.FormValue("code")

		token, err := configs.ConfigGoogle().Exchange(ctx.Context(), code)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		createdUser, err := handler.GoogleService.callback(ctx, token)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.Status(fiber.StatusCreated).JSON(api.Response{
			Data:   createdUser,
			Status: api.StatusSuccess,
		})
	})
}
