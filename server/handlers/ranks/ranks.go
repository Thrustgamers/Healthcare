package rankhandlers

import (
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	return c.SendString("Rank Test")
}
