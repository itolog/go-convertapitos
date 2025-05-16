package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/middleware"
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

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())

	app.Get("/", middleware.Protected(), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// Repositories
	userRepository := user.NewRepository(database)
	// Services
	userService := user.NewService(userRepository)
	// Handlers
	auth.NewHandler(app, auth.HandlerDeps{Config: conf, UserService: userService})
	user.NewHandler(app, user.HandlerDeps{Config: conf, UserServices: userService})

	err := app.Listen(":" + conf.Port)

	if err != nil {
		fmt.Println("App Error", err)
	}
}
