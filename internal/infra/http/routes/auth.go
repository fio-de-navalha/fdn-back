package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupAuthRouter(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", func(c *fiber.Ctx) error {
		return container.AuthHandler.Register(c)
	})

	auth.Post("/login", func(c *fiber.Ctx) error {
		return container.AuthHandler.Login(c)
	})
}
