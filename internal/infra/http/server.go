package http

import (
	"fmt"
	"os"

	"github.com/fio-de-navalha/fdn-back/internal/infra/http/routes"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	app := fiber.New()
	routes.FiberSetupRouters(app)
	app.Listen(":" + os.Getenv("PORT"))
	fmt.Println("Server running...")
}
