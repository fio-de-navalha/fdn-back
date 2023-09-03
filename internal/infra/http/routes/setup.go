package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FiberSetupRouters(app *fiber.App) {
	router := app.Group("/api")

	setupHealthRouter(router)
	setupCustomerRouter(router)
	setupBarberRouter(router)
	setupServiceRouter(router)
	setupProductRouter(router)
}
