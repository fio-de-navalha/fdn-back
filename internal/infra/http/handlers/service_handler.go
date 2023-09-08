package handlers

import (
	"log"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ServiceHandler) Create(c *fiber.Ctx) error {
	log.Println("[handlers.Create] - Validating parameters")
	body := new(service.CreateServiceInput)
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

	input := service.CreateServiceInput{
		BarberId:      body.BarberId,
		Name:          body.Name,
		Price:         body.Price,
		DurationInMin: body.DurationInMin,
	}

	err := h.serviceService.CreateService(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *ServiceHandler) Update(c *fiber.Ctx) error {
	log.Println("[handlers.Update] - Validating parameters")
	serviceId := c.Params("serviceId")
	body := new(service.UpdateServiceInput)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	input := service.UpdateServiceInput{
		Name:          body.Name,
		Price:         body.Price,
		DurationInMin: body.DurationInMin,
		Available:     body.Available,
	}

	err := h.serviceService.UpdateService(serviceId, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Send(nil)
}
