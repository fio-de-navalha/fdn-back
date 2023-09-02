package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupBarberRouter(router fiber.Router) {
	barbers := router.Group("/barbers")
	barbers.Get("/:id", middlewares.EnsureAuth(), container.BarberHandler.GetById)

	auth := router.Group("/auth")
	auth.Post("/register/barbers", container.BarberHandler.Register)
	auth.Post("/login/barbers", container.BarberHandler.Login)
}
