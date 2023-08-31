package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupCustomerRouter(router fiber.Router) {
	customers := router.Group("/customers")

	customers.Get("/:id", func(c *fiber.Ctx) error {
		return container.CustomerHandler.GetUserById(c)
	})
}
