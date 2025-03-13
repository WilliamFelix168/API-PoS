package routes

import (
	"fiber-project/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes sets up all the user routes
func SetupUserRoutes(app *fiber.App) {
	// Group user routes
	api := app.Group("/api")
	users := api.Group("/users")

	// Setup routes
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
}