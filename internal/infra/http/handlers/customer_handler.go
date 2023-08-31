package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
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
