package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/barber"
	"github.com/fio-de-navalha/fdn-back/internal/helpers"
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

func (h *BarberHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.barberService.GetBarberById(id)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if res == nil {
		response := helpers.NewErrorResponse("Barber not found")
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *BarberHandler) Register(c *fiber.Ctx) error {
	body := new(barber.BarberInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	input := barber.BarberInput{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	err := h.barberService.RegisterBarber(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *BarberHandler) Login(c *fiber.Ctx) error {
	body := new(barber.LoginInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	input := barber.LoginInput{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := h.barberService.LoginBarber(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.JSON(res)
}
