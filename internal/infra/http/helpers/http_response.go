package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var errorsMap = map[string]int{
	"permission denied":   fiber.StatusForbidden,
	"invalid credentials": fiber.StatusUnauthorized,
	"not found":           fiber.StatusNotFound,
	"alredy exists":       fiber.StatusUnprocessableEntity,
	"service unavailable": fiber.StatusServiceUnavailable,
}

func BuildErrorResponse(c *fiber.Ctx, errMsg string) error {
	for key, statusCode := range errorsMap {
		if strings.Contains(strings.ToLower(errMsg), key) {
			return c.Status(statusCode).JSON(fiber.Map{
				"error": errMsg,
			})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": errMsg,
	})
}
