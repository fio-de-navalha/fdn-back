package handlers

import (
	"log"
	"strings"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
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
	log.Println("[handlers.GetSalonById] - Validating parameters")
	id := c.Params("id")
	res, err := h.salonService.GetSalonById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonHandler) CraeteSalon(c *fiber.Ctx) error {
	log.Println("[handlers.CraeteSalon] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	body := new(salon.CreateSalonRequest)
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

	res, err := h.salonService.CreateSalon(body.Name, user.ID)
	if err != nil {
		if strings.Contains(err.Error(), "alredy exists") {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *SalonHandler) AddSalonMember(c *fiber.Ctx) error {
	log.Println("[handlers.AddSalonMember] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	id := c.Params("id")
	body := new(salon.AddSalonMemberRequest)
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

	if err := h.salonService.AddSalonMember(id, body.ProfessionalId, body.Role, user.ID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "permission denied") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "alredy exists") {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonHandler) AddSalonAddress(c *fiber.Ctx) error {
	log.Println("[handlers.AddSalonAddress] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	body := new(salon.AddSalonAddressRequest)
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

	if err := h.salonService.AddSalonAddress(user.ID, body.Address); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonHandler) AddSalonContact(c *fiber.Ctx) error {
	log.Println("[handlers.AddSalonContact] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	body := new(salon.AddSalonContactRequest)
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

	if err := h.salonService.AddSalonContact(user.ID, body.Contact); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonHandler) UpdateSalonAddress(c *fiber.Ctx) error {
	log.Println("[handlers.UpdateSalonAddress] - Validating parameters")
	if _, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser); !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	body := new(salon.AddSalonAddressRequest)
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

	res, err := h.salonService.UpdateSalonAddress(salonId, addressId, body.Address)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonHandler) UpdateSalonContact(c *fiber.Ctx) error {
	log.Println("[handlers.UpdateSalonContact] - Validating parameters")
	if _, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser); !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	salonId := c.Params("salonId")
	contactId := c.Params("contactId")
	body := new(salon.AddSalonContactRequest)
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

	res, err := h.salonService.UpdateSalonContact(salonId, contactId, body.Contact)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonHandler) RemoveSalonAddress(c *fiber.Ctx) error {
	log.Println("[handlers.RemoveSalonAddress] - Validating parameters")
	if _, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser); !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	salonId := c.Params("salonId")
	addressId := c.Params("addressId")
	if err := h.salonService.RemoveSalonAddress(salonId, addressId); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *SalonHandler) RemoveSalonContact(c *fiber.Ctx) error {
	log.Println("[handlers.RemoveSalonContact] - Validating parameters")
	if _, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser); !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	salonId := c.Params("salonId")
	contactId := c.Params("contactId")
	if err := h.salonService.RemoveSalonContact(salonId, contactId); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
