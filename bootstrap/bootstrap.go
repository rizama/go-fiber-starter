package bootstrap

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizama/go-fiber-starter/app/middleware"
	"github.com/rizama/go-fiber-starter/config"
	"github.com/rizama/go-fiber-starter/router"
)

func NewApplication() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Starter App v1.0.0",
		ErrorHandler:  middleware.ErrorHandler,
	})
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
		CacheAge: 60,
	}

	app.Use(swagger.New(cfg))

	config.ConnectDb()

	app.Use(idempotency.New())

	app.Use(recover.New())

	app.Use(config.SetupLogger())

	router.Setup(app)

	return app
}
