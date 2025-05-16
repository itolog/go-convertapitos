package cookie

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"time"
)

type Payload struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
	// Allowed session duration
	// Optional. Default value 24 * time.Hour
	Expires time.Duration `json:"expires"`
	// KeyLookup is a string in the form of "<source>:<name>" that is used
	// to extract session id from the request.
	// Possible values: "header:<name>", "query:<name>" or "cookie:<name>"
	// Optional. Default value "cookie:session_id"
	CookieName string `json:"key_lookup"`
}

type Cookie struct {
	cookieDomain string `env:"COOKIE_DOMAIN" env-default:"localhost"`
}

func NewCookie() *Cookie {
	return &Cookie{
		cookieDomain: environments.GetEnv("COOKIE_DOMAIN"),
	}
}

func (cookie *Cookie) SetCookie(c *fiber.Ctx, payload Payload) error {
	sameSite := "lax"
	if environments.IsDev() {
		sameSite = "none"
	}

	sessionStore := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   !environments.IsDev(),
		CookieDomain:   cookie.cookieDomain,
		Expiration:     payload.Expires,
		CookieSameSite: sameSite,
		KeyLookup:      payload.CookieName,
	})

	sess, err := sessionStore.Get(c)
	if err != nil {
		return err
	}

	sess.Set(payload.Key, payload.Value)

	err = sess.Save()
	if err != nil {
		return err
	}

	return nil
}
