package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/authorization"
	"github.com/rs/zerolog/log"

	"github.com/itolog/go-convertapitos/backend/pkg/environments"

	jwtware "github.com/gofiber/contrib/jwt"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(environments.GetEnv("JWT_SECRET"))},
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	if refreshToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, api.ErrUnauthorized)
	}

	authorizationService, authErr := authorization.NewAuthorization()
	if authErr != nil {
		log.Error().Msg(fmt.Sprintf("Authorization Service %v", authErr.Error()))
	}

	_, authErr = authorizationService.VerifyToken(refreshToken)
	if authErr != nil {
		return fiber.NewError(fiber.StatusUnauthorized, authErr.Error())
	}
	return c.Next()
}

func jwtError(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusUnauthorized, api.ErrMissingOrMalformedJWT)
	}

	return fiber.NewError(fiber.StatusUnauthorized, api.ErrInvalidToken)
}
