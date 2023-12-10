package controller

import (
	"log/slog"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/utils"
	"github.com/fio-de-navalha/fdn-back/pkg/validation"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *SalonController) AddSalonAddress(c *fiber.Ctx) error {
	slog.Info("[SalonController.AddSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		slog.Info("[SalonController.AddSalonAddress] - permission denied")
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	if err := validation.ValidUUID(salonId); err != nil {
		slog.Info("[SalonController.AddSalonAddress] - invalid salonId")
		return helpers.BuildErrorResponse(c, err.Error())
	}

	body := new(salon.AddSalonAddressRequest)
	if err := c.BodyParser(&body); err != nil {
		slog.Info("[SalonController.AddSalonAddress] - unable to parse body")
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[SalonController.AddSalonAddress] - Request body: " + utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		slog.Info("[SalonController.UpdateSalonAddress] - request body validation error")
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if err := h.salonService.AddSalonAddress(salonId, user.ID, body.Address); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonController) UpdateSalonAddress(c *fiber.Ctx) error {
	slog.Info("[SalonController.UpdateSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		slog.Info("[SalonController.UpdateSalonAddress] - permission denied")
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	body := new(salon.AddSalonAddressRequest)
	if err := c.BodyParser(&body); err != nil {
		slog.Info("[SalonController.UpdateSalonAddress] - unable to parse body")
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[SalonController.UpdateSalonAddress] - Request body: " + utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		slog.Info("[SalonController.UpdateSalonAddress] - request body validation error")
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.salonService.UpdateSalonAddress(salonId, user.ID, addressId, body.Address)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonController) RemoveSalonAddress(c *fiber.Ctx) error {
	slog.Info("[SalonController.RemoveSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	if err := h.salonService.RemoveSalonAddress(salonId, user.ID, addressId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
