package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupAppointmentRouter(router fiber.Router) {
	appointmentHandler := handlers.NewAppointmentHandler(*container.AppointmentService)

	router.Get("/professional/:professionalId/appointments", appointmentHandler.GetProfessionalAppointments)
	router.Get("/customer/:customerId/appointments", appointmentHandler.GetCustomerAppointments)

	router.Post("/appointment", middlewares.EnsureAuth(), appointmentHandler.Create)
	router.Delete("/appointment/:appointmentId", middlewares.EnsureAuth(), appointmentHandler.Cancel)
}
