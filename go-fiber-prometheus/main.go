package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of requests received",
		},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		requestsTotal.Inc()
		return c.SendString("Welcome to Go API")
	})

	// Expose Prometheus metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	log.Fatal(app.Listen(":3001"))
}
