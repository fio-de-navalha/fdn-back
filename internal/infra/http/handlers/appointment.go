package handlers

import (
	"log"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
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

func (h *AppointmentHandler) GetProfessionalAppointments(c *fiber.Ctx) error {
	log.Println("[handlers.GetProfessionalAppointments] - Validating parameters")
	professionalId := c.Params("professionalId")
	startsAtQuery := c.Query("startsAt")
	if startsAtQuery == "" {
		startsAtQuery = time.Now().Format(constants.DateLayout)
	}

	startsAt, err := time.Parse(constants.DateLayout, startsAtQuery)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.appointmentService.GetProfessionalAppointments(professionalId, startsAt)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentHandler) GetCustomerAppointments(c *fiber.Ctx) error {
	log.Println("[handlers.GetCustomerAppointments] - Validating parameters")
	id := c.Params("customerId")
	res, err := h.appointmentService.GetCustomerAppointments(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentHandler) GetAppointment(c *fiber.Ctx) error {
	log.Println("[handlers.GetAppointment] - Validating parameters")
	id := c.Params("id")
	res, err := h.appointmentService.GetAppointment(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	log.Println("[handlers.Create] - Validating parameters")
	body := new(appointment.CreateAppointmentRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}
	if user.ID != body.ProfessionalId && user.ID != body.CustomerId {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	if body.StartsAt.Before(time.Now()) {
		return helpers.BuildErrorResponse(c, "cannot create appointment in the past")
	}

	input := appointment.CreateAppointmentRequest{
		ProfessionalId: body.ProfessionalId,
		CustomerId:     body.CustomerId,
		StartsAt:       body.StartsAt,
		ServiceIds:     body.ServiceIds,
		ProductIds:     body.ProductIds,
	}
	err := h.appointmentService.CreateApppointment(input)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *AppointmentHandler) Cancel(c *fiber.Ctx) error {
	log.Println("[handlers.Cancel] - Validating parameters")
	appointmentId := c.Params("appointmentId")

	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	err := h.appointmentService.CancelApppointment(user.ID, appointmentId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).Send(nil)
}
