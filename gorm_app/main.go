package main

import (
	"gormapp/config"
	"gormapp/model"
	"gormapp/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	// config.DB.AutoMigrate(&model.User{}) //it ll create table in database automatically

	//config.DB.AutoMigrate(&model.User{}, &model.Profile{})

	config.DB.AutoMigrate(&model.User{}, &model.Profile{}, &model.Order{})

	router.SetupRoutes(app)

	app.Listen(":8080")
}
