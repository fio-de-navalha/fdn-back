package handlers

import (
	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
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

func (h *ProductHandler) GetByBarberId(c *fiber.Ctx) error {
	barberId := c.Params("barberId")

	res, err := h.productService.GetProductsByBarberId(barberId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	body := new(product.CreateProductInput)
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

	input := product.CreateProductInput{
		BarberId: body.BarberId,
		Name:     body.Name,
		Price:    body.Price,
	}

	err := h.productService.CreateProduct(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	param := struct {
		ProductId uint `params:"productId"`
	}{}
	c.ParamsParser(&param)

	body := new(product.UpdateProductInput)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	input := product.UpdateProductInput{
		Name:      body.Name,
		Price:     body.Price,
		Available: body.Available,
	}

	err := h.productService.UpdateProduct(param.ProductId, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Send(nil)
}