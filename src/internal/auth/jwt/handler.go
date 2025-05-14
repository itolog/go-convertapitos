package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
)

type HandlerDeps struct {
	*configs.Config
}

type HandlerJwtAuth struct {
	*configs.Config
}

func NewJWTAuthHandler(router fiber.Router, deps HandlerDeps) {
	handler := HandlerJwtAuth{
		Config: deps.Config,
	}

	router.Post("/login", handler.login)
	router.Post("/register", handler.register)
}
