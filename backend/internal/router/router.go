package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/role"
	_ "github.com/itolog/go-convertapitos/docs"
	"github.com/rs/zerolog"

	"github.com/itolog/go-convertapitos/backend/configs"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/auth"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
	"github.com/itolog/go-convertapitos/backend/internal/home"
	"github.com/itolog/go-convertapitos/backend/pkg/db"
)

type Deps struct {
	*configs.Config
	Database     *db.Db
	CustomLogger *zerolog.Logger
}

type Routes struct {
	*configs.Config
	Database *db.Db
}

// @title			ConvertApiTos API
// @version		1.0.0
// @description	The ConvertApiTos API
// @BasePath		/api/v1
func New(app fiber.Router, deps Deps) {
	apiV1 := app.Group("api/v1")
	// Repositories
	userRepository := user.NewRepository(deps.Database)
	roleRepository := role.NewRepository(deps.Database)
	// Services
	userService := user.NewService(userRepository)
	roleService := role.NewService(roleRepository)
	// Handlers
	home.NewHandler(app)

	apiV1.Get("/swagger/*", swagger.New(swagger.Config{
		Title: "ConvertApiTos API",
		SyntaxHighlight: &swagger.SyntaxHighlightConfig{
			Theme:    "monokai",
			Activate: true,
		},
	}))

	auth.NewAuthHandler(apiV1, auth.Deps{
		Config:       deps.Config,
		UserService:  userService,
		CustomLogger: deps.CustomLogger,
	})
	user.NewHandler(apiV1, user.HandlerDeps{Config: deps.Config, UserServices: userService})

	role.NewHandler(apiV1, role.HandlerDeps{Config: deps.Config, RoleServices: roleService})
}
