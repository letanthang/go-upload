package route

import (
	"github.com/labstack/echo"
	"github.com/letanthang/go-upload/handler/public"
)

func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")
	publicRoute.GET("/health", handler_public.HealthCheck)
	publicRoute.GET("/hello", handler_public.Hello)
	publicRoute.GET("/high", handler_public.High)
	publicRoute.GET("/low", handler_public.Low)
	publicRoute.POST("/upload", handler_public.Upload)
}
