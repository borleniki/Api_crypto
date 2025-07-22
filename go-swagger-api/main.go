package main

import (
	_ "go-swagger-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Rajesh Swagger API
// @version 1.0
// @description This is Fiber REST API with Swagger
// @host localhost:3000
// @BasePath /

// @contact.name Rajesh API Support
// @contact.url https://rkcptraining.com
// @contact.email upadhyay.rajesh@rediffmail.com

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API with Swagger!")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // Swagger endpoint

	app.Get("/hello", HelloWorld) // Sample documented endpoint

	app.Listen(":3000")
}

// HelloWorld godoc
// @Summary Greet the world
// @Description Say Hello
// @Tags Example
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, World!"
// @Router /hello [get]
func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
