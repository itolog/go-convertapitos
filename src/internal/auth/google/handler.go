package google

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/itolog/go-convertapitos/src/configs"
	"time"
)

type HandlerDeps struct {
	*configs.Config
}

type HandlerGoogleAuth struct {
	*configs.Config
	SessionStore *session.Store
}

func NewGoogleAuthHandler(router fiber.Router, deps HandlerDeps) {
	sameSite := "lax"
	if configs.IsDev() {
		sameSite = "none"
	}

	sessionStore := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   !configs.IsDev(),
		CookieDomain:   deps.Auth.CookieDomain,
		CookieSameSite: sameSite,
	})

	handler := HandlerGoogleAuth{
		Config:       deps.Config,
		SessionStore: sessionStore,
	}

	router.Get("/google", handler.login)
	router.Get("/google/callback", handler.callback)
}

type CookiePayload struct {
	Token     string        `json:"token"`
	Value     any           `json:"value"`
	Expires   time.Duration `json:"expires"`
	KeyLookup string        `json:"key_lookup"`
}

func (handler *HandlerGoogleAuth) setCookie(c *fiber.Ctx, payload CookiePayload) error {
	sameSite := "lax"
	if configs.IsDev() {
		sameSite = "none"
	}

	sessionStore := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   !configs.IsDev(),
		CookieDomain:   handler.Auth.CookieDomain,
		Expiration:     payload.Expires,
		CookieSameSite: sameSite,
		KeyLookup:      payload.KeyLookup,
	})

	sess, err := sessionStore.Get(c)
	if err != nil {
		return err
	}

	sess.Set(payload.Token, payload.Value)

	err = sess.Save()
	if err != nil {
		return err
	}

	return nil
}
