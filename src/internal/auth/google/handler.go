package google

import (
	"github.com/gofiber/fiber/v2"
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

	router.Get("/google", handler.GoogleService.login)
	router.Get("/google/callback", handler.GoogleService.callback)
}
