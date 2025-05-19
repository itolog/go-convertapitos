package authorization

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"github.com/itolog/go-convertapitos/src/pkg/jwt"
	"time"
)

type Authorization struct {
	JWT         *jwt.JWT
	cookieStore *session.Store
}

const CookieKey = "refreshToken"

func NewAuthorization() (*Authorization, error) {
	jwtService, err := jwt.NewJWT()
	if err != nil {
		return nil, err
	}

	return &Authorization{
		JWT: jwtService,
	}, nil
}

func (auth *Authorization) SetAuth(ctx *fiber.Ctx, email string) (*string, error) {
	tokens, err := auth.JWT.GenAccessTokens(jwt.Payload{Email: email})
	if err != nil {
		return nil, err
	}

	auth.SetCookie(ctx, tokens.RefreshToken, auth.JWT.RefreshTokenExpires)

	return &tokens.AccessToken, nil
}

func (auth *Authorization) SetCookie(ctx *fiber.Ctx, token string, expires time.Duration) {
	sameSite := "lax"
	if environments.IsDev() {
		sameSite = "none"
	}

	cookie := new(fiber.Cookie)
	cookie.Name = CookieKey
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(expires)
	cookie.SameSite = sameSite
	cookie.Domain = environments.GetEnv("COOKIE_DOMAIN")
	cookie.Secure = !environments.IsDev()

	ctx.Cookie(cookie)
}

func (auth *Authorization) VerifyToken(refreshToken string) (*jwt.UserClaims, error) {
	verify, err := auth.JWT.Verify(refreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return verify, nil
}
