package handlers

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *SalonHandler) AddSalonAddress(c *fiber.Ctx) error {
	log.Println("[SalonHandler.AddSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	body := new(salon.AddSalonAddressRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonHandler.AddSalonAddress] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if err := h.salonService.AddSalonAddress(user.ID, body.Address); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonHandler) UpdateSalonAddress(c *fiber.Ctx) error {
	log.Println("[SalonHandler.UpdateSalonAddress] - Validating parameters")
	if _, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser); !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	body := new(salon.AddSalonAddressRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonHandler.UpdateSalonAddress] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.salonService.UpdateSalonAddress(salonId, addressId, body.Address)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonHandler) RemoveSalonAddress(c *fiber.Ctx) error {
	log.Println("[SalonHandler.RemoveSalonAddress] - Validating parameters")
	if _, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser); !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	if err := h.salonService.RemoveSalonAddress(salonId, addressId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
