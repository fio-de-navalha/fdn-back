package routes

import (
	"github.com/fio-de-navalha/fdn-back/internal/infra/container"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/handlers"
	"github.com/fio-de-navalha/fdn-back/internal/infra/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func setupProfessionalRouter(router fiber.Router) {
	professionalHandler := handlers.NewProfessionalHandler(*container.ProfessionalService)

	professionals := router.Group("/professional")
	professionals.Get("/:id", professionalHandler.GetProfessionalById)

	auth := router.Group("/auth")
	auth.Post("/register/professional", professionalHandler.RegisterProfessional)
	auth.Post("/login/professional", professionalHandler.LoginProfessional)
	auth.Get("/me/professional", middlewares.EnsureProfessionalRole(), professionalHandler.MeProfessional)
}
