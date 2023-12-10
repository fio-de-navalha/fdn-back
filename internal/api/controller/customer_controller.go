package controller

import (
	"log/slog"
	"strings"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/customer"
	"github.com/fio-de-navalha/fdn-back/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	customerService         app.CustomerService
	verificationCodeService app.VerificationCodeService
}

func NewCustomerController(
	customerService app.CustomerService,
	verificationCodeService app.VerificationCodeService,
) *CustomerController {
	return &CustomerController{
		customerService:         customerService,
		verificationCodeService: verificationCodeService,
	}
}

func (h *CustomerController) GetCustomerById(c *fiber.Ctx) error {
	slog.Info("[CustomerController.GetCustomerById] - Validating parameters")
	id := c.Params("id")
	res, err := h.customerService.GetCustomerById(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *CustomerController) RegisterCustomer(c *fiber.Ctx) error {
	slog.Info("[CustomerController.RegisterCustomer] - Validating parameters")
	body := new(customer.RegisterRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[CustomerController.RegisterCustomer] - Request body: " + utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input := customer.RegisterRequest{
		Name:     body.Name,
		Phone:    body.Phone,
		Password: body.Password,
		Question: body.Question,
		Answer:   body.Answer,
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
	slog.Info("[CustomerController.LoginCustomer] - Validating parameters")
	body := new(customer.LoginRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[CustomerController.LoginCustomer] - Request body: " + utils.StructStringfy(&body))
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
	slog.Info("[CustomerController.MeCustomer] - Validating token")
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

func (h *CustomerController) ForgotPassword(c *fiber.Ctx) error {
	slog.Info("[CustomerController.ForgotPassword] - Validating parameters")
	body := new(customer.ForgotPasswordRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[CustomerController.ForgotPassword] - Request body: " + utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input := customer.ForgotPasswordRequest{
		Phone:    body.Phone,
		Question: body.Question,
		Answer:   body.Answer,
	}
	cus, err := h.customerService.ForgotPassword(input)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	token := h.verificationCodeService.GenerateTemporaryToken(cus.ID, cus.Phone)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *CustomerController) UpdatePassword(c *fiber.Ctx) error {
	slog.Info("[CustomerController.UpdatePassword] - Validating parameters")
	body := new(customer.UpdatePasswordRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[CustomerController.UpdatePassword] - Request body: " + utils.StructStringfy(&body))
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	token := strings.Split(body.Token, "-")
	decodedPhone, err := utils.Base64Decode(token[1])
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	cus, err := h.customerService.GetCustomerByPhone(decodedPhone)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	if _, isValid := h.verificationCodeService.GetCacheRegistry(cus.ID + ":" + "token"); !isValid {
		return helpers.BuildErrorResponse(c, "invalid temporary token")
	}

	if _, err := h.customerService.UpdateCustomerPassword(decodedPhone, body.Password); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
