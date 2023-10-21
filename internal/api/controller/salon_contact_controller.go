package controller

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/utils"
	"github.com/fio-de-navalha/fdn-back/pkg/validation"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *SalonController) AddSalonContact(c *fiber.Ctx) error {
	log.Println("[SalonController.AddSalonContact] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	if err := validation.ValidUUID(salonId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	body := new(salon.AddSalonContactRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonController.AddSalonContact] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if err := h.salonService.AddSalonContact(salonId, user.ID, body.Contact); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonController) UpdateSalonContact(c *fiber.Ctx) error {
	log.Println("[SalonController.UpdateSalonContact] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	contactId := c.Params("contactId")
	body := new(salon.AddSalonContactRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[SalonController.UpdateSalonContact] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := h.salonService.UpdateSalonContact(salonId, user.ID, contactId, body.Contact)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonController) RemoveSalonContact(c *fiber.Ctx) error {
	log.Println("[SalonController.RemoveSalonContact] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	contactId := c.Params("contactId")
	if err := h.salonService.RemoveSalonContact(salonId, user.ID, contactId); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
