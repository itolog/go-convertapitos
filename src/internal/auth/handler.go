package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth/google"
	"github.com/itolog/go-convertapitos/src/internal/auth/jwt"
	"github.com/itolog/go-convertapitos/src/internal/user"
)

type HandlerDeps struct {
	*configs.Config
	UserRepository *user.Repository
}

func NewHandler(app *fiber.App, deps HandlerDeps) {
	router := app.Group("/auth")
	// Services
	jwtService := jwt.NewService(jwt.ServiceDeps{
		UserRepository: deps.UserRepository,
		Config:         deps.Config,
	})
	googleService := google.NewService(google.ServiceDeps{
		UserRepository: deps.UserRepository,
		Config:         deps.Config,
	})
	// Handlers
	jwt.NewHandler(router, jwt.HandlerDeps{
		JwtService: jwtService,
	})
	google.NewHandler(router, google.HandlerDeps{
		GoogleService: googleService,
	})

}
