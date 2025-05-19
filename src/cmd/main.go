package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/router"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/db"
	"github.com/itolog/go-convertapitos/src/pkg/logger"
)

func main() {
	conf := configs.NewConfig()
	database := db.NewDb(conf)
	logConfig := configs.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New(fiber.Config{
		Prefork:      true,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: api.ErrorHandler,
		AppName:      "ConvertApiTos",
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	app.Static("/public", "./src/public")

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	router.New(app, router.Deps{
		Config:       conf,
		Database:     database,
		CustomLogger: customLogger,
	})
	err := app.Listen(":" + conf.Port)

	if err != nil {
		customLogger.Error().Msg(fmt.Sprintf("Server error %v", err.Error()))
	}
}
