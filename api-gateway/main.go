package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	// Recovery middleware to prevent crashes
	app.Use(recover.New())

	// Route to service1
	app.All("/service1/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://service1:3000")
	})
	app.All("/service2/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://service2:3000")
	})

	// Start API Gateway
	app.Listen(":8080")
}