package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupProductRouter(router fiber.Router) {
	productHandler := handlers.NewProductHandler(*container.ProductService)

	products := router.Group("/salon/:salonId")

	products.Get("/products", productHandler.GetBySalonId)
	products.Post("/products", middlewares.EnsureProfessionalRole(), productHandler.Create)
	products.Put("/products/:productId", middlewares.EnsureProfessionalRole(), productHandler.Update)
}
