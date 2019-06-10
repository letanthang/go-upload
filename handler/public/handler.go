package handler_public

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world! I love you so much. :3:3:3:4")
}
