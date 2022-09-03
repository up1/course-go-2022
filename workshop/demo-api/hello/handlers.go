package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

type HelloRequest struct {
	Message string `json:"message"`
}

type HelloResponse struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

var tracer = otel.Tracer("echo-server")

func HelloPostHandler(c echo.Context) error {
	_, span := tracer.Start(c.Request().Context(), "HelloPostHandler")
	defer span.End()

	h := new(HelloRequest)
	if err := c.Bind(h); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate input
	// Call service
	todoNext(c.Request().Context())

	// Create and send response to requester
	res := HelloResponse{Message: h.Message, Code: 200}
	return c.JSON(http.StatusOK, res)
}

func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloResponse{})
}
