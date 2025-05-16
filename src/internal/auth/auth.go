package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth/google"
	"github.com/itolog/go-convertapitos/src/internal/user"
)

type Deps struct {
	*configs.Config
	UserService *user.Service
}

func NewAuthHandler(app *fiber.App, deps Deps) {
	router := app.Group("/auth")
	// JWT Auth
	authService := NewService(ServiceDeps{
		UserService: deps.UserService,
	})
	NewHandler(router, HandlerDeps{
		AuthService: authService,
	})
	// Google Auth
	googleService := google.NewService(google.ServiceDeps{
		UserService: deps.UserService,
	})
	google.NewHandler(router, google.HandlerDeps{
		GoogleService: googleService,
	})
}
