package loginhandlers

import (
	"api/utils"

	"github.com/gofiber/fiber/v2"
)

func StatusCheck(c *fiber.Ctx) error {

	return utils.SendSuccessResponse(c, "")
}
