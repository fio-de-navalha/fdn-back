package api

import (
	"log"
	"strconv"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type SalonProductController struct {
	productService application.ProductService
}

func NewSalonProductController(productService application.ProductService) *SalonProductController {
	return &SalonProductController{
		productService: productService,
	}
}

func (h *SalonProductController) GetBySalonId(c *fiber.Ctx) error {
	log.Println("[ProductController.GetBySalonId] - Validating parameters")
	salonId := c.Params("salonId")
	res, err := h.productService.GetProductsBySalonId(salonId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *SalonProductController) Create(c *fiber.Ctx) error {
	log.Println("[ProductController.Create] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	price, _ := strconv.Atoi(c.FormValue("price"))
	input := salon.CreateProductRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
		Name:           c.FormValue("name"),
		Price:          price,
	}

	log.Println("[ProductController.Create] - Request body:", utils.StructStringfy(&input))
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	file, _ := c.FormFile("file")
	err := h.productService.CreateProduct(input, file)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *SalonProductController) Update(c *fiber.Ctx) error {
	log.Println("[ProductController.Update] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	productId := c.Params("productId")
	input := salon.UpdateProductRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
	}

	log.Println("[ProductController.Update] - Request body:", utils.StructStringfy(&input))
	if name := c.FormValue("name"); name != "" {
		input.Name = &name
	}
	if priceStr := c.FormValue("price"); priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err == nil {
			input.Price = &price
		}
	}
	if availableStr := c.FormValue("available"); availableStr != "" {
		available, err := strconv.ParseBool(availableStr)
		if err == nil {
			input.Available = &available
		}
	}

	file, _ := c.FormFile("file")
	err := h.productService.UpdateProduct(productId, input, file)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Send(nil)
}
