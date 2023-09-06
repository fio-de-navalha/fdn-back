package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentService application.AppointmentService
}

func NewAppointmentHandler(appointmentService application.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

func (h *AppointmentHandler) GetBarberAppointments(c *fiber.Ctx) error {
	id := c.Params("barberId")
	res, err := h.appointmentService.GetBarberAppointments(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentHandler) GetCustomerAppointments(c *fiber.Ctx) error {
	id := c.Params("customerId")
	res, err := h.appointmentService.GetCustomerAppointments(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentHandler) GetAppointment(c *fiber.Ctx) error {
	param := struct {
		Id uint `params:"id"`
	}{}
	c.ParamsParser(&param)
	res, err := h.appointmentService.GetAppointment(param.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	body := new(appointment.CreateAppointmentRequest)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	input := appointment.CreateAppointmentRequest{
		BarberId:   body.BarberId,
		CustomerId: body.CustomerId,
		StartsAt:   body.StartsAt,
		ServiceIds: body.ServiceIds,
		ProductIds: body.ProductIds,
	}
	err := h.appointmentService.CreateApppointment(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}