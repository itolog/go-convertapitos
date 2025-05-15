package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/db"
)

func main() {
	conf := configs.NewConfig()
	database := db.NewDb(conf)

	app := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// Handlers
	auth.NewHandler(app, auth.HandlerDeps{Config: conf})
	user.NewHandler(app, user.HandlerDeps{Config: conf, Database: database})

	err := app.Listen(":" + conf.Port)

	if err != nil {
		fmt.Println("App Error", err)
	}
}
