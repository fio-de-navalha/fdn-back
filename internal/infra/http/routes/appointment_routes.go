package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupAppointmentRouter(router fiber.Router) {
	appointmentHandler := api.NewAppointmentHandler(*container.AppointmentService)

	router.Get("/professional/:professionalId/appointments", appointmentHandler.GetProfessionalAppointments)
	router.Get("/customer/:customerId/appointments", appointmentHandler.GetCustomerAppointments)

	router.Post("/appointment", middlewares.EnsureAuth(), appointmentHandler.Create)
	router.Delete("/appointment/:appointmentId", middlewares.EnsureAuth(), appointmentHandler.Cancel)
}
