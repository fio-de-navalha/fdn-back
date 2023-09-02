package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupCustomerRouter(router fiber.Router) {
	customers := router.Group("/customers")
	customers.Get("/:id", middlewares.EnsureAuth(), container.CustomerHandler.GetById)

	auth := router.Group("/auth")
	auth.Post("/register/customers", container.CustomerHandler.Register)
	auth.Post("/login/customers", container.CustomerHandler.Login)
}
