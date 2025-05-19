package home

import (
	"github.com/gofiber/fiber/v2"
	templadaptor "github.com/itolog/go-convertapitos/src/pkg/templ-adaptor"
	"github.com/itolog/go-convertapitos/src/views"
)

type Handler struct{}

func NewHandler(app *fiber.App) {
	handler := &Handler{}

	app.Get("/", handler.Index)
}

func (h *Handler) Index(c *fiber.Ctx) error {
	component := views.Main()
	return templadaptor.Render(c, component)
}
