package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/service"
	"github.com/fio-de-navalha/fdn-back/internal/helpers"
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
	barberId := c.Params("barberId")

	res, err := h.serviceService.GetServicesByBarberId(barberId)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ServiceHandler) Create(c *fiber.Ctx) error {
	body := new(service.CreateServiceDto)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	input := service.CreateServiceDto{
		BarberId:      body.BarberId,
		Name:          body.Name,
		Price:         body.Price,
		DurationInMin: body.DurationInMin,
	}

	err := h.serviceService.CreateService(input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *ServiceHandler) Update(c *fiber.Ctx) error {
	serviceId := c.Params("serviceId")
	body := new(service.UpdateServiceDto)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	input := service.UpdateServiceDto{
		Name:          body.Name,
		Price:         body.Price,
		DurationInMin: body.DurationInMin,
		IsAvailable:   body.IsAvailable,
	}

	err := h.serviceService.UpdateService(serviceId, input)
	if err != nil {
		response := helpers.NewErrorResponse(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.JSON(nil)
}
