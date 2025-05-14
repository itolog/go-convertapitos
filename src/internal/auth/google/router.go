package google

import (
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	handlers := newAuthGoogleHandler()

	router.Get("/google", handlers.auth)
	router.Get("/google/callback", handlers.callback)
	router.Get("/google/profile", handlers.profile)
}
