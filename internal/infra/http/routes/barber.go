package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupBarberRouter(router fiber.Router) {
	barbers := router.Group("/barber")
	barbers.Get("/:id", container.BarberHandler.GetBarberById)

	auth := router.Group("/auth")
	auth.Post("/register/barber", container.BarberHandler.RegisterBarber)
	auth.Post("/login/barber", container.BarberHandler.LoginBarber)
	auth.Get("/me/barber", middlewares.EnsureBarberRole(), container.BarberHandler.MeBarber)
}
