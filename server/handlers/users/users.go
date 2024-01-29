package userhandlers

import (
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	return c.SendString("User Test")
}
