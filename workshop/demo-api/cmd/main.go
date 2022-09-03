package main

import (
	"api/hello"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)
	e.GET("/", hello.HelloHandler)
	e.POST("/hello", hello.HelloPostHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
