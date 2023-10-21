package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api/controller"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupSalonServiceRouter(router fiber.Router) {
	serviceService := container.LoadServiceService()
	serviceController := controller.NewSalonServiceController(*serviceService)

	services := router.Group("/salon/:salonId")

	services.Get("/services", serviceController.GetBySalonId)
	services.Post("/services", middlewares.EnsureProfessionalRole(), serviceController.Create)
	services.Put("/services/:serviceId", middlewares.EnsureProfessionalRole(), serviceController.Update)
}
