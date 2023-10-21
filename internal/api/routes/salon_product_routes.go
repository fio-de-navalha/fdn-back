package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api/controller"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupSalonProductRouter(router fiber.Router) {
	productController := controller.NewSalonProductController(*container.ProductService)

	products := router.Group("/salon/:salonId")

	products.Get("/products", productController.GetBySalonId)
	products.Post("/products", middlewares.EnsureProfessionalRole(), productController.Create)
	products.Put("/products/:productId", middlewares.EnsureProfessionalRole(), productController.Update)
}
