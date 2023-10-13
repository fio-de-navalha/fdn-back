package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupServiceRouter(router fiber.Router) {
	serviceController := api.NewServiceController(*container.ServiceService)

	services := router.Group("/salon/:salonId")

	services.Get("/services", serviceController.GetBySalonId)
	services.Post("/services", middlewares.EnsureProfessionalRole(), serviceController.Create)
	services.Put("/services/:serviceId", middlewares.EnsureProfessionalRole(), serviceController.Update)
}
