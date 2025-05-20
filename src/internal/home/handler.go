package home

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func NewHandler(app fiber.Router) {
	handler := &Handler{}

	app.Get("/", handler.Index)
}

func (h *Handler) Index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
