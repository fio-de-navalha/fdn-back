package api

import (
	"log"
	"strconv"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ServiceHandler struct {
	serviceService application.ServiceService
}

func NewServiceHandler(serviceService application.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		serviceService: serviceService,
	}
}

func (h *ServiceHandler) GetBySalonId(c *fiber.Ctx) error {
	log.Println("[ServiceHandler.GetBySalonId] - Validating parameters")
	salonId := c.Params("salonId")
	res, err := h.serviceService.GetServicesBySalonId(salonId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ServiceHandler) Create(c *fiber.Ctx) error {
	log.Println("[ServiceHandler.Create] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "Permission denied")
	}

	salonId := c.Params("salonId")
	price, _ := strconv.Atoi(c.FormValue("price"))
	durationInMin, _ := strconv.Atoi(c.FormValue("durationInMin"))
	input := service.CreateServiceRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
		Name:           c.FormValue("name"),
		Description:    c.FormValue("description"),
		Price:          price,
		DurationInMin:  durationInMin,
	}

	log.Println("[ServiceHandler.Create] - Request body:", utils.StructStringfy(&input))
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	file, _ := c.FormFile("file")
	if err := h.serviceService.CreateService(input, file); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *ServiceHandler) Update(c *fiber.Ctx) error {
	log.Println("[ServiceHandler.Update] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "Permission denied")
	}

	salonId := c.Params("salonId")
	serviceId := c.Params("serviceId")

	input := service.UpdateServiceRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
	}

	log.Println("[ServiceHandler.Update] - Request body:", utils.StructStringfy(&input))
	if name := c.FormValue("name"); name != "" {
		input.Name = &name
	}
	if description := c.FormValue("description"); description != "" {
		input.Description = &description
	}
	if priceStr := c.FormValue("price"); priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err == nil {
			input.Price = &price
		}
	}
	if durationInMinStr := c.FormValue("durationInMin"); durationInMinStr != "" {
		durationInMin, err := strconv.Atoi(durationInMinStr)
		if err == nil {
			input.DurationInMin = &durationInMin
		}
	}
	if availableStr := c.FormValue("available"); availableStr != "" {
		available, err := strconv.ParseBool(availableStr)
		if err == nil {
			input.Available = &available
		}
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	file, _ := c.FormFile("file")
	err := h.serviceService.UpdateService(serviceId, input, file)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Send(nil)
}
