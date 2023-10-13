package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupAppointmentRouter(router fiber.Router) {
	appointmentController := api.NewAppointmentController(*container.AppointmentService)

	router.Get("/professional/:professionalId/appointments", appointmentController.GetProfessionalAppointments)
	router.Get("/customer/:customerId/appointments", appointmentController.GetCustomerAppointments)

	router.Post("/appointment", middlewares.EnsureAuth(), appointmentController.Create)
	router.Delete("/appointment/:appointmentId", middlewares.EnsureAuth(), appointmentController.Cancel)
}
