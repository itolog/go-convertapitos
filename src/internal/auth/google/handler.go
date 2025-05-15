package google

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"time"
)

type HandlerDeps struct {
	*configs.Config
}

type Handler struct {
	*configs.Config
	SessionStore *session.Store
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	//sameSite := "lax"
	//if environments.IsDev() {
	//	sameSite = "none"
	//}
	//
	//sessionStore := session.New(session.Config{
	//	CookieHTTPOnly: true,
	//	CookieSecure:   !environments.IsDev(),
	//	CookieDomain:   deps.Auth.CookieDomain,
	//	CookieSameSite: sameSite,
	//})

	handler := Handler{
		Config: deps.Config,
		//SessionStore: sessionStore,
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

func (handler *Handler) setCookie(c *fiber.Ctx, payload CookiePayload) error {
	sameSite := "lax"
	if environments.IsDev() {
		sameSite = "none"
	}

	sessionStore := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   !environments.IsDev(),
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
