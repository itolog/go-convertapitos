package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	googleAuth "github.com/itolog/go-convertapitos/src/internal/auth/google"
	"github.com/itolog/go-convertapitos/src/pkg/config"
)

func main() {
	conf := config.NewConfig()

	app := fiber.New(fiber.Config{
		Prefork: true,
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

	googleAuth.Routes(app)

	err := app.Listen(":" + conf.Port)

	if err != nil {
		fmt.Println("App Error", err)
	}
}
