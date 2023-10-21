package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/api/controller"
	"github.com/fio-de-navalha/fdn-back/internal/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupProfessionalRouter(router fiber.Router) {
	professionalService := container.LoadProfessionalService()
	professionalController := controller.NewProfessionalController(*professionalService)

	professionals := router.Group("/professional")
	professionals.Get("/:id", professionalController.GetProfessionalById)

	auth := router.Group("/auth")
	auth.Post("/register/professional", professionalController.RegisterProfessional)
	auth.Post("/login/professional", professionalController.LoginProfessional)
	auth.Get("/me/professional", middlewares.EnsureProfessionalRole(), professionalController.MeProfessional)
}
