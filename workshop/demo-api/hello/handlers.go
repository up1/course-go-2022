package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}