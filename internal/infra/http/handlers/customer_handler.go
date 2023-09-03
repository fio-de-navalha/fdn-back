package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService application.CustomerService
}

func NewCustomerHandler(customerService application.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.customerService.GetCustomerById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if res == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CustomerHandler) Register(c *fiber.Ctx) error {
	body := new(customer.RegisterRequest)
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

	input := customer.RegisterRequest{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
	}

	err := h.customerService.RegisterCustomer(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *CustomerHandler) Login(c *fiber.Ctx) error {
	body := new(customer.LoginRequest)
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

	input := customer.LoginRequest{
		Phone:    body.Phone,
		Password: body.Password,
	}

	resp, err := h.customerService.LoginCustomer(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(resp)
}
