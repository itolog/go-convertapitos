package authorization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/cookie"
	"github.com/itolog/go-convertapitos/src/pkg/jwt"
	"time"
)

type Authorization struct {
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func (auth *Authorization) SetAuth(ctx *fiber.Ctx, email string) (*string, error) {
	jwtService, err := jwt.NewJWT()
	if err != nil {
		return nil, err
	}

	tokens, err := jwtService.GenAccessTokens(email)
	if err != nil {
		return nil, err
	}

	err = auth.SetCookie(ctx, tokens.RefreshToken, jwtService.RefreshTokenExpires)
	if err != nil {
		return nil, err
	}

	return &tokens.AccessToken, nil
}

func (auth *Authorization) SetCookie(ctx *fiber.Ctx, token string, expires time.Duration) error {
	cookieStore := cookie.NewCookie()
	err := cookieStore.SetCookie(ctx, cookie.Payload{
		Key:        "refreshToken",
		Value:      token,
		CookieName: "cookie:refreshToken",
		Expires:    expires,
	})
	if err != nil {
		return err
	}
	return nil
}
