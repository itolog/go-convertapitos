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
	"github.com/gofiber/swagger"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/auth"
	"github.com/itolog/go-convertapitos/src/internal/home"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/db"
	"github.com/itolog/go-convertapitos/src/pkg/logger"

	_ "github.com/itolog/go-convertapitos/docs"
)

// @title			ConvertApiTos API
// @version		1.0.0
// @description	The ConvertApiTos API
// @BasePath		/
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

	apiV1 := app.Group("api/v1")

	// Repositories
	userRepository := user.NewRepository(database)
	// Services
	userService := user.NewService(userRepository)
	// Handlers
	apiV1.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	home.NewHandler(app)
	auth.NewAuthHandler(apiV1, auth.Deps{
		Config:       conf,
		UserService:  userService,
		CustomLogger: customLogger,
	})
	user.NewHandler(apiV1, user.HandlerDeps{Config: conf, UserServices: userService})

	err := app.Listen(":" + conf.Port)

	if err != nil {
		customLogger.Error().Msg(fmt.Sprintf("Server error %v", err.Error()))
	}
}
