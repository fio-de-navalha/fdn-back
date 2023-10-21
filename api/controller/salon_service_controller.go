package controller

import (
	"log"
	"strconv"

	"github.com/fio-de-navalha/fdn-back/api/helpers"
	"github.com/fio-de-navalha/fdn-back/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/app"
	"github.com/fio-de-navalha/fdn-back/constants"
	"github.com/fio-de-navalha/fdn-back/domain/salon"
	"github.com/fio-de-navalha/fdn-back/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type SalonServiceController struct {
	serviceService app.ServiceService
}

func NewSalonServiceController(serviceService app.ServiceService) *SalonServiceController {
	return &SalonServiceController{
		serviceService: serviceService,
	}
}

func (h *SalonServiceController) GetBySalonId(c *fiber.Ctx) error {
	log.Println("[ServiceController.GetBySalonId] - Validating parameters")
	salonId := c.Params("salonId")
	res, err := h.serviceService.GetServicesBySalonId(salonId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonServiceController) Create(c *fiber.Ctx) error {
	log.Println("[ServiceController.Create] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "Permission denied")
	}

	salonId := c.Params("salonId")
	price, _ := strconv.Atoi(c.FormValue("price"))
	durationInMin, _ := strconv.Atoi(c.FormValue("durationInMin"))
	input := salon.CreateServiceRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
		Name:           c.FormValue("name"),
		Description:    c.FormValue("description"),
		Price:          price,
		DurationInMin:  durationInMin,
	}

	log.Println("[ServiceController.Create] - Request body:", utils.StructStringfy(&input))
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

func (h *SalonServiceController) Update(c *fiber.Ctx) error {
	log.Println("[ServiceController.Update] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "Permission denied")
	}

	salonId := c.Params("salonId")
	serviceId := c.Params("serviceId")

	input := salon.UpdateServiceRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
	}

	log.Println("[ServiceController.Update] - Request body:", utils.StructStringfy(&input))
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
