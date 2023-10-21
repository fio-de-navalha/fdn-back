package controller

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *SalonController) AddSalonAddress(c *fiber.Ctx) error {
	log.Println("[SalonController.AddSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	if err := utils.ValidUUID(salonId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	body := new(salon.AddSalonAddressRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonController.AddSalonAddress] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if err := h.salonService.AddSalonAddress(salonId, user.ID, body.Address); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonController) UpdateSalonAddress(c *fiber.Ctx) error {
	log.Println("[SalonController.UpdateSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	body := new(salon.AddSalonAddressRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonController.UpdateSalonAddress] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.salonService.UpdateSalonAddress(salonId, user.ID, addressId, body.Address)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonController) RemoveSalonAddress(c *fiber.Ctx) error {
	log.Println("[SalonController.RemoveSalonAddress] - Validating parameters")
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
