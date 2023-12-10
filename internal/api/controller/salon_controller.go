package controller

import (
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type SalonController struct {
	salonService app.SalonService
}

func NewSalonController(salonService app.SalonService) *SalonController {
	return &SalonController{
		salonService: salonService,
	}
}

func (h *SalonController) GetSalonById(c *fiber.Ctx) error {
	slog.Info("[SalonController.GetSalonById] - Validating parameters")
	salonId := c.Params("salonId")
	res, err := h.salonService.GetSalonById(salonId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonController) CraeteSalon(c *fiber.Ctx) error {
	slog.Info("[SalonController.CraeteSalon] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	body := new(salon.CreateSalonRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[SalonController.CraeteSalon] - Request body: " + utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.salonService.CreateSalon(body.Name, user.ID)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
