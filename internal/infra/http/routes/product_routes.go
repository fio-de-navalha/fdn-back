package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupProductRouter(router fiber.Router) {
	productHandler := api.NewProductHandler(*container.ProductService)

	products := router.Group("/salon/:salonId")

	products.Get("/products", productHandler.GetBySalonId)
	products.Post("/products", middlewares.EnsureProfessionalRole(), productHandler.Create)
	products.Put("/products/:productId", middlewares.EnsureProfessionalRole(), productHandler.Update)
}
