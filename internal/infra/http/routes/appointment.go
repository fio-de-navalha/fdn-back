package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupAppointmentRouter(router fiber.Router) {
	router.Get("/barber/:barberId/appointments", container.AppointmentHandler.GetBarberAppointments)
	router.Get("/customer/:customerId/appointments", container.AppointmentHandler.GetCustomerAppointments)
}
