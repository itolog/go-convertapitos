package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth/google"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/authorization"
)

type Deps struct {
	*configs.Config
	UserService user.IUserService
}

func NewAuthHandler(app *fiber.App, deps Deps) {
	router := app.Group("/auth")
	authorizationService, err := authorization.NewAuthorization()
	if err != nil {
		fmt.Println("Authorization Service", err.Error())
	}
	// JWT Auth
	authService := NewService(ServiceDeps{
		UserService:   deps.UserService,
		Authorization: authorizationService,
	})
	NewHandler(router, HandlerDeps{
		AuthService: authService,
	})
	// Google Auth
	googleService := google.NewService(google.ServiceDeps{
		UserService:   deps.UserService,
		Authorization: authorizationService,
	})
	google.NewHandler(router, google.HandlerDeps{
		GoogleService: googleService,
	})
}
