package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Example route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Service 2!")
	})

	// Start server
	app.Listen(":3000")
}