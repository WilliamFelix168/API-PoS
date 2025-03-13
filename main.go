package main

import (
	"log"

	"fiber-project/config"
	"fiber-project/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize Fiber with config
	app := fiber.New(config.FiberConfig())

	// Middleware
	app.Use(logger.New())   // Log HTTP requests
	app.Use(recover.New())  // Recover from panics
	app.Use(cors.New())     // Enable CORS

	// Setup routes
	routes.SetupUserRoutes(app)

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the User API",
			"version": "1.0.0",
		})
	})

	// Handle 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Route not found",
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}