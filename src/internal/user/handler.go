package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/db"
)

type HandlerDeps struct {
	*configs.Config
	Database *db.Db
}

type Handler struct {
	*configs.Config
	repository *Repository
}

func NewHandler(app *fiber.App, deps HandlerDeps) {
	repository := NewRepository(deps.Database)
	router := app.Group("/user")

	handler := Handler{
		Config:     deps.Config,
		repository: repository,
	}

	router.Get("/", handler.findAll)
	router.Get("/:id", handler.findById)
	router.Post("/", handler.create)
	router.Patch("/:id", handler.update)
	router.Delete("/:id", handler.delete)
}
