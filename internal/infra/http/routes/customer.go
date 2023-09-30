package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/labstack/echo/v4"
)

func setupCustomerRouter(r *echo.Group) {
	customerHandler := handlers.NewCustomerHandler(*container.CustomerService)

	customers := r.Group("/customer")
	customers.GET("/:id", authMiddleware, customerHandler.GetCustomerById)

	auth := r.Group("/auth")
	auth.POST("/register/customer", customerHandler.RegisterCustomer)
	auth.POST("/login/customer", customerHandler.LoginCustomer)
	auth.GET("/me/customer", middlewares.EnsureAuth(), customerHandler.MeCustomer)
}
