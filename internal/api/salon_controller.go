package api

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type SalonHandler struct {
	salonService application.SalonService
}

func NewSalonHandler(salonService application.SalonService) *SalonHandler {
	return &SalonHandler{
		salonService: salonService,
	}
}

func (h *SalonHandler) GetSalonById(c *fiber.Ctx) error {
	log.Println("[SalonHandler.GetSalonById] - Validating parameters")
	salonId := c.Params("salonId")
	res, err := h.salonService.GetSalonById(salonId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonHandler) CraeteSalon(c *fiber.Ctx) error {
	log.Println("[SalonHandler.CraeteSalon] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	body := new(salon.CreateSalonRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonHandler.CraeteSalon] - Request body:", utils.StructStringfy(&body))
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
