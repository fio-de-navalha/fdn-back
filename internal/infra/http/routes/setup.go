package routes

import "github.com/labstack/echo/v4"

func FiberSetupRouters(e *echo.Echo) {
	router := e.Group("/api")

	setupHealthRouter(router)
	setupCustomerRouter(router)
	setupProfessionalRouter(router)
	setupSalonRouter(router)
	setupServiceRouter(router)
	setupProductRouter(router)
	setupAppointmentRouter(router)
}
