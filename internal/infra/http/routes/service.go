package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupServiceRouter(router fiber.Router) {
	services := router.Group("/barbers/:barberId/services")
	services.Get("/:serviceId", middlewares.EnsureAuth(), container.CustomerHandler.GetUserById)
}
