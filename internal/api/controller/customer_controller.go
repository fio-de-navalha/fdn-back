package controller

import (
	"log"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	customerService app.CustomerService
}

func NewCustomerController(customerService app.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (h *CustomerController) GetCustomerById(c *fiber.Ctx) error {
	log.Println("[CustomerController.GetCustomerById] - Validating parameters")
	id := c.Params("id")
	res, err := h.customerService.GetCustomerById(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CustomerController) RegisterCustomer(c *fiber.Ctx) error {
	log.Println("[CustomerController.RegisterCustomer] - Validating parameters")
	body := new(customer.RegisterRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[CustomerController.RegisterCustomer] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input := customer.RegisterRequest{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
	}

	res, err := h.customerService.RegisterCustomer(input)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CustomerController) LoginCustomer(c *fiber.Ctx) error {
	log.Println("[CustomerController.LoginCustomer] - Validating parameters")
	body := new(customer.LoginRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	log.Println("[CustomerController.LoginCustomer] - Request body:", utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input := customer.LoginRequest{
		Phone:    body.Phone,
		Password: body.Password,
	}
	res, err := h.customerService.LoginCustomer(input)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.JSON(res)
}

func (h *CustomerController) MeCustomer(c *fiber.Ctx) error {
	log.Println("[CustomerController.MeCustomer] - Validating token")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}
	res, err := h.customerService.GetCustomerById(user.ID)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
