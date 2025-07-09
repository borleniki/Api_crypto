package router

import (
	"transactionapp/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/accounts")
	api.Get("/", controller.GetAccounts)
	api.Get("/:id", controller.GetAccount)
	api.Post("/", controller.CreateAccount)
	api.Put("/:id", controller.UpdateAccount)
	api.Delete("/:id", controller.DeleteAccount)
	api.Post("/transfer", controller.TransferAmount)
	api.Get("/:id/mini-statement", controller.MiniStatement)
}
