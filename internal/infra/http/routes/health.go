package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupHealthRouter(r *echo.Group) {
	r.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
