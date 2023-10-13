package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupProductRouter(router fiber.Router) {
	productController := api.NewProductController(*container.ProductService)

	products := router.Group("/salon/:salonId")

	products.Get("/products", productController.GetBySalonId)
	products.Post("/products", middlewares.EnsureProfessionalRole(), productController.Create)
	products.Put("/products/:productId", middlewares.EnsureProfessionalRole(), productController.Update)
}
