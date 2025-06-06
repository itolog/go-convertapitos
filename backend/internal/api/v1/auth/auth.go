package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/configs"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
	"github.com/itolog/go-convertapitos/backend/pkg/authorization"
	"github.com/rs/zerolog"
)

type Deps struct {
	*configs.Config
	UserService  user.IUserService
	CustomLogger *zerolog.Logger
}

func NewAuthHandler(app fiber.Router, deps Deps) {
	router := app.Group("/auth")
	authorizationService, err := authorization.NewAuthorization()
	if err != nil {
		deps.CustomLogger.Error().Msg(fmt.Sprintf("Authorization Service %v", err.Error()))
	}

	authService := NewService(ServiceDeps{
		UserService:   deps.UserService,
		Authorization: authorizationService,
		CustomLogger:  deps.CustomLogger,
	})
	NewHandler(router, HandlerDeps{
		AuthService:  authService,
		CustomLogger: deps.CustomLogger,
	})
}
