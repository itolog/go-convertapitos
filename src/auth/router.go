package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/auth/google", Auth)
	api.Get("/auth/google/callback", Callback)
}
