package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupBarberRouter(router fiber.Router) {
	barbers := router.Group("/barber")
	barbers.Get("/:id", container.BarberHandler.GetBarberById)

	barbers.Post("/:id/address", middlewares.EnsureBarberRole(), container.BarberHandler.AddBarberAddress)
	barbers.Put("/:id/address/:addressId", middlewares.EnsureBarberRole(), container.BarberHandler.UpdateBarberAddress)
	barbers.Delete("/:id/address/:addressId", middlewares.EnsureBarberRole(), container.BarberHandler.RemoveBarberAddress)

	barbers.Post("/:id/contact", middlewares.EnsureBarberRole(), container.BarberHandler.AddBarberContact)
	barbers.Put("/:id/contact/:contactId", middlewares.EnsureBarberRole(), container.BarberHandler.UpdateBarberContact)
	barbers.Delete("/:id/contact/:contactId", middlewares.EnsureBarberRole(), container.BarberHandler.RemoveBarberContact)

	auth := router.Group("/auth")
	auth.Post("/register/barber", container.BarberHandler.RegisterBarber)
	auth.Post("/login/barber", container.BarberHandler.LoginBarber)
	auth.Get("/me/barber", middlewares.EnsureBarberRole(), container.BarberHandler.MeBarber)
}
