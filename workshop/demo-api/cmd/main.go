package main

import (
	"api/hello"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello.HelloHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
