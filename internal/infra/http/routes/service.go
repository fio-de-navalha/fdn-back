package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupServiceRouter(router fiber.Router) {
	serviceHandler := handlers.NewServiceHandler(*container.ServiceService)

	services := router.Group("/barber/:barberId")

	services.Get("/services", serviceHandler.GetByBarberId)
	services.Post("/services", middlewares.EnsureBarberRole(), serviceHandler.Create)
	services.Put("/services/:serviceId", middlewares.EnsureBarberRole(), serviceHandler.Update)
}
