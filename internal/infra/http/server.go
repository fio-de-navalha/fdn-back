package http

import (
	"os"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/infra/http/routes"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Server() {
	app := fiber.New()

	app.Use(idempotency.New())
	app.Use(limiter.New(limiter.Config{
		Max:               15,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${ip}:${port} | ${latency} | ${status} | ${method} ${path}\n",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Backend Metrics Page"}))

	routes.FiberSetupRouters(app)

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: "./api/swagger.json",
	}))

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
