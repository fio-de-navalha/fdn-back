package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FiberSetupRouters(app *fiber.App) {
	router := app.Group("/api")

	setupAppRouter(router)
	setupCustomerRouter(router)
	setupBarberRouter(router)
	setupServiceRouter(router)
}
