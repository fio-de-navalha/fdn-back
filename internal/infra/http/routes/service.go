package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupServiceRouter(router fiber.Router) {
	services := router.Group("/barbers/:barberId")

	services.Get("/services", container.ServiceHandler.GetByBarberId)
	services.Post("/services", middlewares.EnsureBarberRole(), container.ServiceHandler.Create)
	services.Put("/services/:serviceId", middlewares.EnsureBarberRole(), container.ServiceHandler.Update)
}
