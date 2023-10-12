package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupCustomerRouter(router fiber.Router) {
	customerHandler := api.NewCustomerHandler(*container.CustomerService)

	customers := router.Group("/customer")
	customers.Get("/:id", middlewares.EnsureAuth(), customerHandler.GetCustomerById)

	auth := router.Group("/auth")
	auth.Post("/register/customer", customerHandler.RegisterCustomer)
	auth.Post("/login/customer", customerHandler.LoginCustomer)
	auth.Get("/me/customer", middlewares.EnsureAuth(), customerHandler.MeCustomer)
}
