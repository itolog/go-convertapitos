package jwt

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerDeps struct {
	JwtService *Service
}

type Handler struct {
	JwtService *Service
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := Handler{

		JwtService: deps.JwtService,
	}

	router.Post("/login", handler.JwtService.login)
	router.Post("/register", handler.JwtService.register)
}
