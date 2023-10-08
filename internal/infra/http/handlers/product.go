package handlers

import (
	"log"
	"strconv"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/helpers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService application.ProductService
}

func NewProductHandler(productService application.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetBySalonId(c *fiber.Ctx) error {
	log.Println("[ProductHandler.GetBySalonId] - Validating parameters")
	salonId := c.Params("salonId")
	res, err := h.productService.GetProductsBySalonId(salonId)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	log.Println("[ProductHandler.Create] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	price, _ := strconv.Atoi(c.FormValue("price"))
	input := product.CreateProductRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
		Name:           c.FormValue("name"),
		Price:          price,
	}

	log.Println("[ProductHandler.Create] - Request body:", utils.StructPrettify(&input))
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

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	log.Println("[ProductHandler.Update] - Validating parameters")
	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return helpers.BuildErrorResponse(c, "permission denied")
	}

	salonId := c.Params("salonId")
	productId := c.Params("productId")
	input := product.UpdateProductRequest{
		SalonId:        salonId,
		ProfessionalId: user.ID,
	}

	log.Println("[ProductHandler.Update] - Request body:", utils.StructPrettify(&input))
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
