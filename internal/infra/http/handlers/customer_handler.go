package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	customer "github.com/fio-de-navalha/fdn-back/internal/domain/customer/entities"
	"github.com/fio-de-navalha/fdn-back/internal/helpers"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService application.CustomerServices
}

func NewCustomerHandler(customerService application.CustomerServices) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.customerService.GetCustomerById(id)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if user == nil {
		response := helpers.NewErrorResponse("User not found")
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *CustomerHandler) Register(c *fiber.Ctx) error {
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

func (h *CustomerHandler) Login(c *fiber.Ctx) error {
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
