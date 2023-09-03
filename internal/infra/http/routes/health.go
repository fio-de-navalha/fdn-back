package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupHealthRouter(router fiber.Router) {
	router.Get("/health", handlers.GetHealth)
}
