package google

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/auth/google", GoogleAuth)
	api.Get("/auth/google/callback", GoogleAuthCallback)
	api.Get("/auth/google/profile", GoogleProfile)
}
