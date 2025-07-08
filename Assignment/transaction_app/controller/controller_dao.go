package controller

import (
	"transactionapp/model"
	"transactionapp/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAccounts(c *fiber.Ctx) error {
	accounts, err := service.GetAccounts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(accounts)
}
func GetAccount(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	account, err := service.GetAccount(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(account)
}
func CreateAccount(c *fiber.Ctx) error {
	var account model.Account
	if err := c.BodyParser(&account); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := service.Create(account); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "account created"})
}
func UpdateAccount(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var account model.Account
	if err := c.BodyParser(&account); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := service.Update(id, account); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Account updated"})
}
func DeleteAccount(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Account deleted"})
}