package middlewares

import (
	"strings"

	"github.com/fio-de-navalha/fdn-back/pkg/cryptography"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired JWT",
		})
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
		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing JWT Token",
			})
		}
		token := strings.Split(authorization, "Bearer ")
		if len(token) == 1 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing JWT Token",
			})
		}
		jwtToken, err := cryptography.ParseToken(token[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized"},
			)
		}
		if jwtToken["role"] != "barber" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Permission denied"},
			)
		}

		return c.Next()
	}
}
