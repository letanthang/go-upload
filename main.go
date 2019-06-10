package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/letanthang/go-upload/config"
	"github.com/letanthang/go-upload/profiler"
	"github.com/letanthang/go-upload/route"
)

func main() {
	////////////////
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, //1KB
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	if config.Config.Profiler.StatsdAddress != "" {
		e.Use(profiler.ProfilerWithConfig(profiler.ProfilerConfig{Address: config.Config.Profiler.StatsdAddress, Service: config.Config.Profiler.Service}))
	}

	e.File("/form", "form.html")
	route.Public(e)

	fmt.Println("Server listening at 9090")

	port := "9090"
	err := e.Start(":" + port)
	if err != nil {
		fmt.Println(err)
	}
	// http.HandleFunc("/upload", uploadHandler)
	// http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/form", formHandler)
	// if err := http.ListenAndServe(":9090", nil); err != nil {
	// 	panic(err)
	// }
}
