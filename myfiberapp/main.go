package main

import (
	"myfiberapp/config"
	"myfiberapp/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":3034")
}
