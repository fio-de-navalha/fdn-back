package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupBarberRouter(router fiber.Router) {
	barberHandler := handlers.NewBarberHandler(*container.ProfessionalService)

	barbers := router.Group("/barber")
	barbers.Get("/:id", barberHandler.GetBarberById)

	barbers.Post("/:id/address", middlewares.EnsureBarberRole(), barberHandler.AddBarberAddress)
	barbers.Put("/:id/address/:addressId", middlewares.EnsureBarberRole(), barberHandler.UpdateBarberAddress)
	barbers.Delete("/:id/address/:addressId", middlewares.EnsureBarberRole(), barberHandler.RemoveBarberAddress)

	barbers.Post("/:id/contact", middlewares.EnsureBarberRole(), barberHandler.AddBarberContact)
	barbers.Put("/:id/contact/:contactId", middlewares.EnsureBarberRole(), barberHandler.UpdateBarberContact)
	barbers.Delete("/:id/contact/:contactId", middlewares.EnsureBarberRole(), barberHandler.RemoveBarberContact)

	auth := router.Group("/auth")
	auth.Post("/register/barber", barberHandler.RegisterBarber)
	auth.Post("/login/barber", barberHandler.LoginBarber)
	auth.Get("/me/barber", middlewares.EnsureBarberRole(), barberHandler.MeBarber)
}
