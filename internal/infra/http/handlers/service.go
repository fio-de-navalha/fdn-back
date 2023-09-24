package handlers

import (
	"log"
	"strconv"
	"strings"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
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

func (h *ServiceHandler) GetByBarberId(c *fiber.Ctx) error {
	log.Println("[handlers.GetByBarberId] - Validating parameters")
	barberId := c.Params("barberId")
	res, err := h.serviceService.GetServicesByBarberId(barberId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ServiceHandler) Create(c *fiber.Ctx) error {
	log.Println("[handlers.Create] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	price, _ := strconv.Atoi(c.FormValue("price"))
	durationInMin, _ := strconv.Atoi(c.FormValue("durationInMin"))
	input := service.CreateServiceRequest{
		BarberId:      user.ID,
		Name:          c.FormValue("name"),
		Description:   c.FormValue("description"),
		Price:         price,
		DurationInMin: durationInMin,
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	file, _ := c.FormFile("file")
	if err := h.serviceService.CreateService(input, file); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *ServiceHandler) Update(c *fiber.Ctx) error {
	log.Println("[handlers.Update] - Validating parameters")
	serviceId := c.Params("serviceId")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}
	if user.ID != c.Params("barberId") {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	input := service.UpdateServiceRequest{}
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	file, _ := c.FormFile("file")
	err := h.serviceService.UpdateService(serviceId, input, file)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Send(nil)
}
