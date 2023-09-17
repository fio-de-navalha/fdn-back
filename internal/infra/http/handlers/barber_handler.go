package handlers

import (
	"log"
	"strings"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type BarberHandler struct {
	barberService application.BarberService
}

func NewBarberHandler(barberService application.BarberService) *BarberHandler {
	return &BarberHandler{
		barberService: barberService,
	}
}

func (h *BarberHandler) GetBarberById(c *fiber.Ctx) error {
	log.Println("[handlers.GetBarberById] - Validating parameters")
	id := c.Params("id")
	res, err := h.barberService.GetBarberById(id)
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

func (h *BarberHandler) RegisterBarber(c *fiber.Ctx) error {
	log.Println("[handlers.RegisterBarber] - Validating parameters")
	body := new(barber.RegisterRequest)
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

	input := barber.RegisterRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	resp, err := h.barberService.RegisterBarber(input)
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

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *BarberHandler) LoginBarber(c *fiber.Ctx) error {
	log.Println("[handlers.LoginBarber] - Validating parameters")
	body := new(barber.LoginRequest)
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

	input := barber.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := h.barberService.LoginBarber(input)
	if err != nil {
		if err.Error() == "invalid credentials" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(res)
}

func (h *BarberHandler) MeBarber(c *fiber.Ctx) error {
	log.Println("[handlers.MeBarber] - Validating barber")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}
	res, err := h.barberService.GetBarberById(user.ID)
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
