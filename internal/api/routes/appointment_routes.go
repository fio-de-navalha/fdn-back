package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api/controller"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupAppointmentRouter(router fiber.Router) {
	appointmentService := container.LoadAppointmentService()
	appointmentController := controller.NewAppointmentController(*appointmentService)

	router.Get("/professional/:professionalId/appointments", appointmentController.GetProfessionalAppointments)
	router.Get("/customer/:customerId/appointments", appointmentController.GetCustomerAppointments)

	router.Post("/appointment", middlewares.EnsureAuth(), appointmentController.Create)
	router.Delete("/appointment/:appointmentId", middlewares.EnsureAuth(), appointmentController.Cancel)
}
