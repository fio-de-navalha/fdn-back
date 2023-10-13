package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupSalonRouter(router fiber.Router) {
	salonController := api.NewSalonController(*container.SalonService)

	salons := router.Group("/salon")
	salons.Get("/:salonId", salonController.GetSalonById)
	salons.Post("/", middlewares.EnsureProfessionalRole(), salonController.CraeteSalon)
	salons.Post("/:id/members", middlewares.EnsureProfessionalRole(), salonController.AddSalonMember)

	salons.Post("/:salonId/address", middlewares.EnsureProfessionalRole(), salonController.AddSalonAddress)
	salons.Put("/:salonId/address/:addressId", middlewares.EnsureProfessionalRole(), salonController.UpdateSalonAddress)
	salons.Delete("/:salonId/address/:addressId", middlewares.EnsureProfessionalRole(), salonController.RemoveSalonAddress)

	salons.Post("/:salonId/contact", middlewares.EnsureProfessionalRole(), salonController.AddSalonContact)
	salons.Put("/:salonId/contact/:contactId", middlewares.EnsureProfessionalRole(), salonController.UpdateSalonContact)
	salons.Delete("/:salonId/contact/:contactId", middlewares.EnsureProfessionalRole(), salonController.RemoveSalonContact)

	salons.Post("/:salonId/period", middlewares.EnsureProfessionalRole(), salonController.AddSalonPeriod)
	salons.Put("/:salonId/period/:periodId", middlewares.EnsureProfessionalRole(), salonController.UpdateSalonPeriod)
	salons.Delete("/:salonId/period/:periodId", middlewares.EnsureProfessionalRole(), salonController.RemoveSalonPeriod)
}
