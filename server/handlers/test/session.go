package testhandlers

import (
	"api/storage"
	"api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSessions(c *fiber.Ctx) error {
	return utils.SendSuccessResponse(c, storage.SessionManager)
}
