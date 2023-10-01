package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupSalonRouter(router fiber.Router) {
	salonHandler := handlers.NewSalonHandler(*container.SalonService)

	salons := router.Group("/salon")
	salons.Get("/:id", salonHandler.GetSalonById)
	salons.Post("/", middlewares.EnsureProfessionalRole(), salonHandler.CraeteSalon)
	salons.Post("/:id/members", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonMember)

	salons.Post("/:id/address", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonAddress)
	salons.Put("/:id/address/:addressId", middlewares.EnsureProfessionalRole(), salonHandler.UpdateSalonAddress)
	salons.Delete("/:id/address/:addressId", middlewares.EnsureProfessionalRole(), salonHandler.RemoveSalonAddress)

	salons.Post("/:id/contact", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonContact)
	salons.Put("/:id/contact/:contactId", middlewares.EnsureProfessionalRole(), salonHandler.UpdateSalonContact)
	salons.Delete("/:id/contact/:contactId", middlewares.EnsureProfessionalRole(), salonHandler.RemoveSalonContact)

	salons.Post("/:salonId/period", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonPeriod)
	salons.Put("/:salonId/period/:periodId", middlewares.EnsureProfessionalRole(), salonHandler.UpdateSalonPeriod)
	salons.Delete("/:salonId/period/:periodId", middlewares.EnsureProfessionalRole(), salonHandler.RemoveSalonPeriod)
}
