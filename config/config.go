package config

import "github.com/gofiber/fiber/v2"

// FiberConfig returns the Fiber configuration
func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      "Simple User API",
		ErrorHandler: defaultErrorHandler,
	}
}

// defaultErrorHandler is the default error handler
func defaultErrorHandler(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Return JSON error message
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}