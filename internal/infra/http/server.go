package http

import (
	"os"

	"github.com/fio-de-navalha/fdn-back/internal/infra/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Server() {
	app := fiber.New()
	app.Use(logger.New())

	routes.FiberSetupRouters(app)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
