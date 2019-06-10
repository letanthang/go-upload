package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
