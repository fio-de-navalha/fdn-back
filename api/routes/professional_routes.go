package routes

import (
	"github.com/fio-de-navalha/fdn-back/api/controller"
	"github.com/fio-de-navalha/fdn-back/api/middlewares"
	"github.com/fio-de-navalha/fdn-back/infra/container"
	"github.com/gofiber/fiber/v2"
)

func setupProfessionalRouter(router fiber.Router) {
	professionalController := controller.NewProfessionalController(*container.ProfessionalService)

	professionals := router.Group("/professional")
	professionals.Get("/:id", professionalController.GetProfessionalById)

	auth := router.Group("/auth")
	auth.Post("/register/professional", professionalController.RegisterProfessional)
	auth.Post("/login/professional", professionalController.LoginProfessional)
	auth.Get("/me/professional", middlewares.EnsureProfessionalRole(), professionalController.MeProfessional)
}
