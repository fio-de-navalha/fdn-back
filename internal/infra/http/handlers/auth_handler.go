package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	customer "github.com/fio-de-navalha/fdn-back/internal/domain/customer/entities"
	"github.com/fio-de-navalha/fdn-back/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	customerService application.CustomerServices
}

func NewAuthHandler(customerService application.CustomerServices) *AuthHandler {
	return &AuthHandler{
		customerService: customerService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	u := new(customer.CustomerInput)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	input := customer.CustomerInput{
		Name:     u.Name,
		Phone:    u.Phone,
		Password: u.Password,
	}

	err := h.customerService.RegisterCustomer(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	u := new(customer.LoginInput)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	input := customer.LoginInput{
		Phone:    u.Phone,
		Password: u.Password,
	}

	resp, err := h.customerService.LoginCustomer(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.JSON(resp)
}
