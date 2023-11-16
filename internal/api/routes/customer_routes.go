package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api/controller"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupCustomerRouter(router fiber.Router) {
	customerService := container.LoadCustomerService()
	customerController := controller.NewCustomerController(*customerService)

	customers := router.Group("/customer")
	customers.Get("/:id", middlewares.EnsureAuth(), customerController.GetCustomerById)

	auth := router.Group("/auth")
	auth.Post("/register/customer", customerController.RegisterCustomer)
	auth.Post("/login/customer", customerController.LoginCustomer)
	auth.Post("/forgot/customer", customerController.ForgotPassword)
	auth.Get("/me/customer", middlewares.EnsureAuth(), customerController.MeCustomer)
}
