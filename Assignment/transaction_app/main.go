package main

import (
	"transactionapp/config"
	"transactionapp/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":3004")
}
