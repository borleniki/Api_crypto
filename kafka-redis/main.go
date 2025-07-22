package main

import (
	"kafka-redis/database"
	"kafka-redis/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	// Init connections
	database.InitMySQL()
	database.InitRedis()

	// Routes
	app.Post("/user", handlers.CreateUser)
	app.Get("/user/:id", handlers.GetUser)

	app.Listen(":8080")
}
