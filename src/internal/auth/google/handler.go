package google

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
)

type HandlerDeps struct {
	*configs.Config
}

type HandlerGoogleAuth struct {
	*configs.Config
}

func NewGoogleAuthHandler(router fiber.Router, deps HandlerDeps) {
	handler := HandlerGoogleAuth{
		Config: deps.Config,
	}

	router.Get("/google", handler.login)
	router.Get("/google/callback", handler.callback)
}
