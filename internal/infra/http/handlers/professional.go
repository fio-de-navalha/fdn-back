package handlers

import (
	"log"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ProfessionalHandler struct {
	professionalService application.ProfessionalService
}

func NewProfessionalHandler(professionalService application.ProfessionalService) *ProfessionalHandler {
	return &ProfessionalHandler{
		professionalService: professionalService,
	}
}

func (h *ProfessionalHandler) GetProfessionalById(c *fiber.Ctx) error {
	log.Println("[ProfessionalHandler.GetProfessionalById] - Validating parameters")
	id := c.Params("id")
	res, err := h.professionalService.GetProfessionalById(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ProfessionalHandler) RegisterProfessional(c *fiber.Ctx) error {
	log.Println("[ProfessionalHandler.RegisterProfessional] - Validating parameters")
	body := new(professional.RegisterProfessionalRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input := professional.RegisterProfessionalRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := h.professionalService.RegisterProfessional(input)
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

func (h *ProfessionalHandler) LoginProfessional(c *fiber.Ctx) error {
	log.Println("[ProfessionalHandler.LoginProfessional] - Validating parameters")
	body := new(professional.LoginProfessionalRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input := professional.LoginProfessionalRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := h.professionalService.LoginProfessional(input)
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

func (h *ProfessionalHandler) MeProfessional(c *fiber.Ctx) error {
	log.Println("[ProfessionalHandler.MeProfessional] - Validating barber")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "Permission denied")
	}
	res, err := h.professionalService.GetProfessionalById(user.ID)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
