package routes

import (
	"github.com/fio-de-navalha/fdn-back/api/controller"
	"github.com/fio-de-navalha/fdn-back/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupSalonServiceRouter(router fiber.Router) {
	serviceController := controller.NewSalonServiceController(*container.ServiceService)

	services := router.Group("/salon/:salonId")

	services.Get("/services", serviceController.GetBySalonId)
	services.Post("/services", middlewares.EnsureProfessionalRole(), serviceController.Create)
	services.Put("/services/:serviceId", middlewares.EnsureProfessionalRole(), serviceController.Update)
}
