package routes

import (
	"github.com/gofiber/fiber/v2"
)

func setupHealthRouter(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON("Ok")
	})
}
