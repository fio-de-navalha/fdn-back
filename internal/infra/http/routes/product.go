package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupProductRouter(router fiber.Router) {
	productHandler := handlers.NewProductHandler(*container.ProductService)

	products := router.Group("/barber/:barberId")

	products.Get("/products", productHandler.GetByBarberId)
	products.Post("/products", middlewares.EnsureBarberRole(), productHandler.Create)
	products.Put("/products/:productId", middlewares.EnsureBarberRole(), productHandler.Update)
}
