package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth/google"
	"github.com/itolog/go-convertapitos/src/internal/auth/jwt"
)

type HandlerDeps struct {
	*configs.Config
}

func NewHandler(app *fiber.App, deps HandlerDeps) {
	router := app.Group("/auth")

	jwt.NewHandler(router, jwt.HandlerDeps{
		Config: deps.Config,
	})
	google.NewHandler(router, google.HandlerDeps{
		Config: deps.Config,
	})

}
