package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupProductRouter(router fiber.Router) {
	products := router.Group("/barber/:barberId")

	products.Get("/products", container.ProductHandler.GetByBarberId)
	products.Post("/products", middlewares.EnsureBarberRole(), container.ProductHandler.Create)
	products.Put("/products/:productId", middlewares.EnsureBarberRole(), container.ProductHandler.Update)
}
