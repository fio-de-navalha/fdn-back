package controller

import (
	"log/slog"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/appointment"
	"github.com/fio-de-navalha/fdn-back/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AppointmentController struct {
	appointmentService app.AppointmentService
}

func NewAppointmentController(appointmentService app.AppointmentService) *AppointmentController {
	return &AppointmentController{
		appointmentService: appointmentService,
	}
}

func (h *AppointmentController) GetProfessionalAppointments(c *fiber.Ctx) error {
	slog.Info("[AppointmentController.GetProfessionalAppointments] - Validating parameters")
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

func (h *AppointmentController) GetCustomerAppointments(c *fiber.Ctx) error {
	slog.Info("[AppointmentController.GetCustomerAppointments] - Validating parameters")
	id := c.Params("customerId")
	res, err := h.appointmentService.GetCustomerAppointments(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentController) GetAppointment(c *fiber.Ctx) error {
	slog.Info("[AppointmentController.GetAppointment] - Validating parameters")
	id := c.Params("id")
	res, err := h.appointmentService.GetAppointment(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *AppointmentController) Create(c *fiber.Ctx) error {
	slog.Info("[AppointmentController.Create] - Validating parameters")
	body := new(appointment.CreateAppointmentRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[AppointmentController.Create] - Request body: " + utils.StructStringfy(&body))
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

	body.StartsAt = utils.ConvertToGMTMinus3(body.StartsAt)
	if body.StartsAt.Before(time.Now()) {
		slog.Info("[AppointmentController.Create] - Cannot create appointment in the past")
		slog.Info("[AppointmentController.Create] - StartsAt: " + body.StartsAt.String())
		slog.Info("[AppointmentController.Create] - CurrentTime: " + time.Now().String())
		return helpers.BuildErrorResponse(c, "cannot create appointment in the past")
	}

	input := appointment.CreateAppointmentRequest{
		SalonId:        body.SalonId,
		ProfessionalId: body.ProfessionalId,
		CustomerId:     body.CustomerId,
		StartsAt:       body.StartsAt,
		ServiceIds:     body.ServiceIds,
		ProductIds:     body.ProductIds,
	}
	err := h.appointmentService.CreateApppointment(input)
	if err != nil {
		slog.Error("[AppointmentController.Create] - Appointment creation error: " + err.Error())
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *AppointmentController) Cancel(c *fiber.Ctx) error {
	slog.Info("[AppointmentController.Cancel] - Validating parameters")
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
