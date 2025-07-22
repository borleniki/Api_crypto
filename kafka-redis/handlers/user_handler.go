package handlers

import (
	"encoding/json"
	"fmt"
	"kafka-redis/database"
	"kafka-redis/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	cacheKey := fmt.Sprintf("users:%d", time.Now().Unix())
	// Check Redis cache first
	val, err := database.Rdb.Get(database.Ctx, cacheKey).Result()
	if err == nil {
		var users []models.User
		if err := json.Unmarshal([]byte(val), &users); err == nil {
			fmt.Println("Fetched all users from Redis")
			return c.JSON(users)
		}
	}
	// If not found in Redis, fetch from MySQL
	fmt.Println("Fetching users from MySQL")
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	// Cache the users in Redis
	usersJSON1, _ := json.Marshal(users)
	database.Rdb.Set(database.Ctx, cacheKey, usersJSON1, 0)

	fmt.Println("Fetched all users and cached in Redis")
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	cacheKey := "user:" + id

	// 1. Check Redis
	val, err := database.Rdb.Get(database.Ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		json.Unmarshal([]byte(val), &user)
		fmt.Println(" Fetched from Redis")
		return c.JSON(user)
	}

	// 2. If not found, get from MySQL
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// 3. Cache in Redis
	userJSON, _ := json.Marshal(user)
	database.Rdb.Set(database.Ctx, cacheKey, userJSON, 0)

	fmt.Println("Fetched from MySQL and cached")
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	database.DB.Create(&user)

	// Optional: Invalidate or update cache
	/*cacheKey := "user:" + strconv.Itoa(int(user.ID))
	userJSON, _ := json.Marshal(user)
	database.Rdb.Set(database.Ctx, cacheKey, userJSON, 0)
	fmt.Println("User created and cached in Redis")*/
	return c.Status(201).JSON(user)
}
