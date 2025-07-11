package main

import (
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/consul/api"
)

type Transaction struct {
	ID            int     `json:"id"`
	FromAccountID int     `json:"from_acc"`
	ToAccountID   int     `json:"to_acc"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"created_at"`
}

func registerServiceWithConsul() {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = "producer-1"
	registration.Name = "producer-service"
	registration.Address = "localhost"
	registration.Port = 8080
	registration.Check = &api.AgentServiceCheck{
		HTTP:     "http://localhost:8080/health",
		Interval: "10s",
		Timeout:  "1s",
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(" Registered with Consul")
}

func main() {
	app := fiber.New()
	registerServiceWithConsul()
	app.Post("/api/consumer/send", func(c *fiber.Ctx) error {
		tx := new(Transaction)
		println("tx", tx.Amount, tx.CreatedAt, tx.FromAccountID, tx.ToAccountID)
		if err := c.BodyParser(tx); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(tx).
			Post("http://localhost:3004/accounts")

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to call producer"})
		}

		return c.JSON(fiber.Map{
			"message": "Transaction sent to producer",
			"result":  resp.String(),
		})
	})

	app.Listen(":8081")
}
