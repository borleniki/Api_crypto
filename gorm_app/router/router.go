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

	order := app.Group("/orders")
	order.Get("/", controller.GetOrders)
	order.Get("/:id", controller.GetOrder)
	order.Post("/", controller.CreateOrder)
	order.Put("/:id", controller.UpdateOrder)
	order.Delete("/:id", controller.DeleteOrder)

	profile := app.Group("/profiles")
	profile.Get("/", controller.GetProfiles)
	profile.Get("/:id", controller.GetProfile)
	profile.Post("/", controller.CreateProfile)
	profile.Put("/:id", controller.UpdateProfile)
	profile.Delete("/:id", controller.DeleteProfile)
}
