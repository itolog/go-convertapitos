package main

import (
	"fmt"
	"github.com/itolog/go-convertapitos/src/configs"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/itolog/go-convertapitos/src/internal/auth"
	googleAuth "github.com/itolog/go-convertapitos/src/internal/auth/google"
)

func main() {
	conf := configs.NewConfig()

	app := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	var sameSite string = "lax"
	if configs.IsDev() {
		sameSite = "none"
	}

	googleAuth.SessionStore = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   !configs.IsDev(),
		CookieDomain:   conf.Auth.CookieDomain,
		CookieSameSite: sameSite,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	auth.NewHandler(app, auth.HandlerDeps{Config: conf})

	err := app.Listen(":" + conf.Port)

	if err != nil {
		fmt.Println("App Error", err)
	}
}
