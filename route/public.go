package route

import (
	"github.com/labstack/echo"
	"github.com/letanthang/go-upload/handler/public"
)

func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")
	publicRoute.GET("/health", handler_public.HealthCheck)
}
