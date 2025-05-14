package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/internal/auth/google"
)

func Router(app *fiber.App) {
	router := app.Group("/auth")

	google.Router(router)
}
