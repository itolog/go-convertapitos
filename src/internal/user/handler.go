package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
)

type HandlerDeps struct {
	*configs.Config
	UserServices *Service
}

type Handler struct {
	*configs.Config
	UserServices *Service
}

func NewHandler(app *fiber.App, deps HandlerDeps) {
	router := app.Group("/user")

	handler := Handler{
		Config:       deps.Config,
		UserServices: deps.UserServices,
	}

	router.Get("/", handler.UserServices.findAll)
	router.Get("/:id", handler.UserServices.findById)
	router.Get("/by_email/:email", handler.UserServices.findByEmail)
	router.Post("/", handler.UserServices.create)
	router.Patch("/:id", handler.UserServices.update)
	router.Delete("/:id", handler.UserServices.delete)
}
