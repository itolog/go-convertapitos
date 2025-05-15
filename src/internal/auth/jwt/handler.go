package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
)

type HandlerDeps struct {
	*configs.Config
}

type Handler struct {
	*configs.Config
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := Handler{
		Config: deps.Config,
	}

	router.Post("/login", handler.login)
	router.Post("/register", handler.register)
}
