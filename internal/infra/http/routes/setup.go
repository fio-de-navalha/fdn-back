package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func FiberSetupRouters(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Backend Metrics Page"}))

	router := app.Group("/api")

	setupHealthRouter(router)
	setupCustomerRouter(router)
	setupBarberRouter(router)
	setupServiceRouter(router)
	setupProductRouter(router)
}
