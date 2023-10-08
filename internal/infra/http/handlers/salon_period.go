package handlers

import (
	"log"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/fio-de-navalha/fdn-back/internal/validation"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func hourMinuteFormat(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	expectedFormat := constants.OpenCloseLayout
	_, err := time.Parse(expectedFormat, value)
	return err == nil
}

func (h *SalonHandler) AddSalonPeriod(c *fiber.Ctx) error {
	log.Println("[SalonHandler.AddSalonPeriod] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}
	salonId := c.Params("salonId")
	if err := validation.ValidUUID(salonId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	body := new(salon.AddPeriodRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonHandler.AddSalonPeriod] - Request body:", utils.StructPrettify(&body))
	validate := validator.New()
	validate.RegisterValidation("hourMinuteFormat", hourMinuteFormat)
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if err := h.salonService.AddSalonPeriod(salonId, user.ID, *body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonHandler) UpdateSalonPeriod(c *fiber.Ctx) error {
	log.Println("[SalonHandler.UpdateSalonPeriod] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	if err := validation.ValidUUID(salonId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	periodId := c.Params("periodId")
	if err := validation.ValidUUID(periodId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	body := new(salon.UpdatePeriodRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonHandler.UpdateSalonPeriod] - Request body:", utils.StructPrettify(&body))
	validate := validator.New()
	validate.RegisterValidation("hourMinuteFormat", hourMinuteFormat)
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.salonService.UpdateSalonPeriod(salonId, user.ID, periodId, *body)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonHandler) RemoveSalonPeriod(c *fiber.Ctx) error {
	log.Println("[SalonHandler.RemoveSalonPeriod] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	if err := validation.ValidUUID(salonId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	periodId := c.Params("periodId")
	if err := validation.ValidUUID(periodId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if err := h.salonService.RemoveSalonPeriod(salonId, user.ID, periodId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
