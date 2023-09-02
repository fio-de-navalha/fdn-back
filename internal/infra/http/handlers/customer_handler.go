package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
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

	res, err := h.customerService.GetCustomerById(id)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if res == nil {
		response := helpers.NewErrorResponse("User not found")
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CustomerHandler) Register(c *fiber.Ctx) error {
	body := new(customer.CustomerInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	input := customer.CustomerInput{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
	}

	err := h.customerService.RegisterCustomer(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *CustomerHandler) Login(c *fiber.Ctx) error {
	body := new(customer.LoginInput)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	input := customer.LoginInput{
		Phone:    body.Phone,
		Password: body.Password,
	}

	resp, err := h.customerService.LoginCustomer(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.JSON(resp)
}
