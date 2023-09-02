package middlewares

import (
	"strings"

	"github.com/fio-de-navalha/fdn-back/pkg/cryptography"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

func EnsureAuth() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(cryptography.JwtSecret)},
		ErrorHandler: errorHandler,
	})
}

func EnsureBarberRole() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		token := strings.Split(authorization, "Bearer ")

		jwtToken, err := cryptography.ParseToken(token[1])
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{"error": "Unauthorized"})
		}

		if jwtToken["role"] != "barber" {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{"error": "Permission denied"})
		}

		return c.Next()
	}
}
