package handlers

import (
	"log"
	"strings"

	"github.com/fio-de-navalha/fdn-back/internal/application"
	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
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
	log.Println("[handlers.GetByBarberId] - Validating parameters")
	barberId := c.Params("barberId")
	res, err := h.productService.GetProductsByBarberId(barberId)
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

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	log.Println("[handlers.Create] - Validating parameters")
	body := new(product.CreateProductRequest)
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

	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}
	if user.ID != body.BarberId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	input := product.CreateProductRequest{
		BarberId: body.BarberId,
		Name:     body.Name,
		Price:    body.Price,
	}

	err := h.productService.CreateProduct(input)
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

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	log.Println("[handlers.Update] - Validating parameters")
	barberId := c.Params("barberId")
	productId := c.Params("productId")
	body := new(product.UpdateProductRequest)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, ok := c.Locals(constants.UserContextKey).(middlewares.RquestUser)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}
	if user.ID != barberId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Permission denied",
		})
	}

	input := product.UpdateProductRequest{
		Name:      body.Name,
		Price:     body.Price,
		Available: body.Available,
	}

	err := h.productService.UpdateProduct(productId, input)
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
