package adminhandlers

import (
	"api/security"
	"api/utils"

	"github.com/gofiber/fiber/v2"
)

func getBlockedIps(c *fiber.Ctx) error {

	return utils.SendSuccessResponse(c, security.IpManager)
}

func removeBlockedIp(c *fiber.Ctx) error {

	return utils.SendSuccessResponse(c, "succesfully removed ip")
}
