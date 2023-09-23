package http

import (
	"os"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/infra/http/routes"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(cors.New())
	app.Use(idempotency.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/api/health"
		},
		Max:               15,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${ip}:${port} | ${latency} | ${status} | ${method} ${path}\n\n",
	}))
}

func StartServer() {
	app := fiber.New()

	setupMiddlewares(app)
	routes.FiberSetupRouters(app)

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: "./api/swagger.json",
	}))

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
