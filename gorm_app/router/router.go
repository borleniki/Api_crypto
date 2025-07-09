package router

import (
	"gormapp/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/users")
	api.Get("/", controller.GetUsers)
	api.Get("/:id", controller.GetUser)
	api.Post("/", controller.CreateUser)
	api.Put("/:id", controller.UpdateUser)
	api.Delete("/:id", controller.DeleteUser)
}
