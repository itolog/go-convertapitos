package main

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/itolog/go-convertapitos/src/internal/auth"
	googleAuth "github.com/itolog/go-convertapitos/src/internal/auth/google"
	"github.com/itolog/go-convertapitos/src/pkg/config"
)

func main() {
	conf := config.NewConfig()

	app := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	var sameSite string = "lax"
	if config.IsDev() {
		sameSite = "none"
	}

	googleAuth.SessionStore = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   !config.IsDev(),
		CookieDomain:   conf.CookieDomain,
		CookieSameSite: sameSite,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	auth.Router(app)

	err := app.Listen(":" + conf.Port)

	if err != nil {
		fmt.Println("App Error", err)
	}
}
