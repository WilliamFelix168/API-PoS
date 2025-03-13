package controllers

import (
	"fiber-project/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GetAllUsers returns all users
func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    models.Users,
	})
}

// GetUser returns a specific user by ID
func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid ID",
		})
	}

	for _, user := range models.Users {
		if user.ID == id {
			return c.JSON(fiber.Map{
				"success": true,
				"data":    user,
			})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "User not found",
	})
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Set a new ID (in a real app, this would be handled by the database)
	user.ID = len(models.Users) + 1
	
	// Add to our "database"
	models.Users = append(models.Users, *user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}