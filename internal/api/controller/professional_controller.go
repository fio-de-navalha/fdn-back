package controller

import (
	"log/slog"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/api/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/app"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/professional"
	"github.com/fio-de-navalha/fdn-back/pkg/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ProfessionalController struct {
	professionalService app.ProfessionalService
}

func NewProfessionalController(professionalService app.ProfessionalService) *ProfessionalController {
	return &ProfessionalController{
		professionalService: professionalService,
	}
}

func (h *ProfessionalController) GetProfessionalById(c *fiber.Ctx) error {
	slog.Info("[ProfessionalController.GetProfessionalById] - Validating parameters")
	id := c.Params("id")
	res, err := h.professionalService.GetProfessionalById(id)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ProfessionalController) RegisterProfessional(c *fiber.Ctx) error {
	slog.Info("[ProfessionalController.RegisterProfessional] - Validating parameters")
	body := new(professional.RegisterProfessionalRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[ProfessionalController.RegisterProfessional] - Request body: " + utils.StructStringfy(&body))
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

func (h *ProfessionalController) LoginProfessional(c *fiber.Ctx) error {
	slog.Info("[ProfessionalController.LoginProfessional] - Validating parameters")
	body := new(professional.LoginProfessionalRequest)
	if err := c.BodyParser(&body); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	slog.Info("[ProfessionalController.LoginProfessional] - Request body: " + utils.StructStringfy(&body))
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

func (h *ProfessionalController) MeProfessional(c *fiber.Ctx) error {
	slog.Info("[ProfessionalController.MeProfessional] - Validating professional")
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
