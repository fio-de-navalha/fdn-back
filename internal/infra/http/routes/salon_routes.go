package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupSalonRouter(router fiber.Router) {
	salonHandler := api.NewSalonHandler(*container.SalonService)

	salons := router.Group("/salon")
	salons.Get("/:salonId", salonHandler.GetSalonById)
	salons.Post("/", middlewares.EnsureProfessionalRole(), salonHandler.CraeteSalon)
	salons.Post("/:id/members", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonMember)

	salons.Post("/:salonId/address", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonAddress)
	salons.Put("/:salonId/address/:addressId", middlewares.EnsureProfessionalRole(), salonHandler.UpdateSalonAddress)
	salons.Delete("/:salonId/address/:addressId", middlewares.EnsureProfessionalRole(), salonHandler.RemoveSalonAddress)

	salons.Post("/:salonId/contact", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonContact)
	salons.Put("/:salonId/contact/:contactId", middlewares.EnsureProfessionalRole(), salonHandler.UpdateSalonContact)
	salons.Delete("/:salonId/contact/:contactId", middlewares.EnsureProfessionalRole(), salonHandler.RemoveSalonContact)

	salons.Post("/:salonId/period", middlewares.EnsureProfessionalRole(), salonHandler.AddSalonPeriod)
	salons.Put("/:salonId/period/:periodId", middlewares.EnsureProfessionalRole(), salonHandler.UpdateSalonPeriod)
	salons.Delete("/:salonId/period/:periodId", middlewares.EnsureProfessionalRole(), salonHandler.RemoveSalonPeriod)
}
