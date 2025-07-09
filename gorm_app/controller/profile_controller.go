package controller

import (
	"gormapp/model"
	"gormapp/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProfiles(c *fiber.Ctx) error {
	profiles, err := service.GetProfiles()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(profiles)
}
func GetProfile(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	profile, err := service.GetProfile(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(profile)
}
func CreateProfile(c *fiber.Ctx) error {
	var profile model.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := service.CreateProfile(profile); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Profile created"})
}
func UpdateProfile(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var profile model.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	profile.ID = id
	if err := service.UpdateProfile(profile); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Profile updated"})
}
func DeleteProfile(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteProfile(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Profile deleted"})
}
