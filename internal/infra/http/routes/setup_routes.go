package routes

import (
	"github.com/gofiber/fiber/v2"
)

func FiberSetupRouters(app *fiber.App) {
	router := app.Group("/api")

	setupHealthRouter(router)
	setupCustomerRouter(router)
	setupProfessionalRouter(router)
	setupSalonRouter(router)
	setupSalonServiceRouter(router)
	setupSalonProductRouter(router)
	setupAppointmentRouter(router)
}
