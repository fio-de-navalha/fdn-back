package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupCustomerRouter(router fiber.Router) {
	customers := router.Group("/customer")
	customers.Get("/:id", middlewares.EnsureAuth(), container.CustomerHandler.GetCustomerById)

	auth := router.Group("/auth")
	auth.Post("/register/customer", container.CustomerHandler.RegisterCustomer)
	auth.Post("/login/customer", container.CustomerHandler.LoginCustomer)
	auth.Get("/me/customer", middlewares.EnsureAuth(), container.CustomerHandler.MeCustomer)
}
