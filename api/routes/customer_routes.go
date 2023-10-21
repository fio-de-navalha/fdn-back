package routes

import (
	"github.com/fio-de-navalha/fdn-back/api/controller"
	"github.com/fio-de-navalha/fdn-back/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupCustomerRouter(router fiber.Router) {
	customerController := controller.NewCustomerController(*container.CustomerService)

	customers := router.Group("/customer")
	customers.Get("/:id", middlewares.EnsureAuth(), customerController.GetCustomerById)

	auth := router.Group("/auth")
	auth.Post("/register/customer", customerController.RegisterCustomer)
	auth.Post("/login/customer", customerController.LoginCustomer)
	auth.Get("/me/customer", middlewares.EnsureAuth(), customerController.MeCustomer)
}
