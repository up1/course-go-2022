package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	return e
}
