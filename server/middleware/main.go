package middleware

import (
	"api/storage"

	"github.com/gofiber/fiber/v2"
)

func SessionAuth(c *fiber.Ctx) error {
	// Retrieve session ID from request header or cookies
	sessionID := 1

	if _, ok := storage.SessionManager[sessionID]; ok {
		c.Next()
	}

	// Return unauthorized status if session is not found
	return c.JSON(fiber.Map{
		"status": 404,
		"error":  "Session not found",
	})
}
