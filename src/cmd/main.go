package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/itolog/go-convertapitos/src/auth"
	"github.com/itolog/go-convertapitos/src/config"
)

func main() {
	conf := config.NewConfig()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	auth.Routes(app)

	err := app.Listen(":" + conf.Port)

	if err != nil {
		fmt.Println("App Error", err)
	}
}
