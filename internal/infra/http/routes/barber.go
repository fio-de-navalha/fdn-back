package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupBarberRouter(router fiber.Router) {
	barbers := router.Group("/barber")
	barbers.Get("/:id", container.BarberHandler.GetById)

	auth := router.Group("/auth")
	auth.Post("/register/barber", container.BarberHandler.Register)
	auth.Post("/login/barber", container.BarberHandler.Login)
}
