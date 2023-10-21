package middlewares

import (
	"strings"

	"github.com/fio-de-navalha/fdn-back/constants"
	"github.com/fio-de-navalha/fdn-back/infra/encryption"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RquestUser struct {
	ID string
}

func extractAndValidateToken(c *fiber.Ctx) (jwt.MapClaims, error) {
	authorization := c.Get("Authorization")
	if authorization == "" {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing JWT Token",
		})
	}
	token := strings.Split(authorization, "Bearer ")
	if len(token) == 1 {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing JWT Token",
		})
	}
	jwtToken, err := encryption.ParseToken(token[1])
	if err != nil {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return jwtToken, nil
}

func extractAndSetUser(c *fiber.Ctx, token jwt.MapClaims) error {
	id := token["sub"]
	str, ok := id.(string)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Unable to determine requester user",
		})
	}

	user := RquestUser{
		ID: str,
	}
	c.Locals(constants.UserContextKey, user)
	return nil
}

func EnsureAuth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var jwtToken jwt.MapClaims
		var err error
		if jwtToken, err = extractAndValidateToken(c); err != nil {
			return err
		}
		if err = extractAndSetUser(c, jwtToken); err != nil {
			return err
		}
		return c.Next()
	}
}

func EnsureProfessionalRole() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var jwtToken jwt.MapClaims
		var err error
		if jwtToken, err = extractAndValidateToken(c); err != nil {
			return err
		}
		if jwtToken["role"] != "professional" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Permission denied",
			})
		}
		if err = extractAndSetUser(c, jwtToken); err != nil {
			return err
		}
		return c.Next()
	}
}
