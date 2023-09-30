package middlewares

import (
	"net/http"
	"strings"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/infra/cryptography"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RquestUser struct {
	ID string
}

func extractAndValidateToken(c echo.Context) (jwt.MapClaims, error) {
	authorization := c.Get("Authorization").(string)
	if authorization == "" {
		return nil, c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Missing JWT Token",
		})
	}
	token := strings.Split(authorization, "Bearer ")
	if len(token) == 1 {
		return nil, c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Missing JWT Token",
		})
	}
	jwtToken, err := cryptography.ParseToken(token[1])
	if err != nil {
		return nil, c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Unauthorized",
		})
	}
	return jwtToken, nil
}

func extractAndSetUser(c echo.Context, token jwt.MapClaims) error {
	id := token["sub"]
	str, ok := id.(string)
	if !ok {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": "Unable to determine requester user",
		})
	}

	user := RquestUser{
		ID: str,
	}
	c.Set(constants.UserContextKey, user)
	return nil
}

func EnsureAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var jwtToken jwt.MapClaims
			var err error
			if jwtToken, err = extractAndValidateToken(c); err != nil {
				return err
			}
			if err = extractAndSetUser(c, jwtToken); err != nil {
				return err
			}
			return next(c)
		}
	}
}

func EnsureProfessionalRole() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            var jwtToken jwt.MapClaims
            var err error
            if jwtToken, err = extractAndValidateToken(c); err != nil {
                return err
            }
            if jwtToken["role"] != "professional" {
                return c.JSON(http.StatusForbidden, echo.Map{
                    "error": "Permission denied",
                })
            }
            if err = extractAndSetUser(c, jwtToken); err != nil {
                return err
            }
            return next(c)
        }
    }
}