package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupServiceRouter(router fiber.Router) {
	serviceHandler := handlers.NewServiceHandler(*container.ServiceService)

	services := router.Group("/salon/:salonId")

	services.Get("/services", serviceHandler.GetBySalonId)
	services.Post("/services", middlewares.EnsureProfessionalRole(), serviceHandler.Create)
	services.Put("/services/:serviceId", middlewares.EnsureProfessionalRole(), serviceHandler.Update)
}
