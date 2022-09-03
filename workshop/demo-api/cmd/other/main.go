package main

import (
	"api/hello"
	"api/internal"
	"context"
	"flag"
	"log"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func main() {
	url := flag.String("zipkin", "http://localhost:9411/api/v2/spans", "zipkin url")
	flag.Parse()

	// Init tracer
	tp, err := internal.InitTracer(*url, "other-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Init router
	e := echo.New()
	e.Use(otelecho.Middleware("other-service"))
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.GET("/", hello.HelloHandler)

	e.Logger.Fatal(e.Start(":1324"))
}
