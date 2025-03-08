package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
	
	"billing-service/types"
)


func main() {
	app := fiber.New()

	// Sample route for creating a new bill
	app.Post("/billing", func(c *fiber.Ctx) error {
		billing := new(types.Billing)

		// Parse the incoming JSON payload
		if err := c.BodyParser(billing); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		// Set the timestamp for the billing entry
		billing.Timestamp = time.Now().Format(time.RFC3339)

		// Process payment logic (simulation)
		billing.Status = "Success" // Set to 'Failed' if payment fails in real scenario

		// Return a response
		return c.Status(fiber.StatusOK).JSON(billing)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
