package adminhandlers

import (
	"api/security"
	"api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetBlockedIps(c *fiber.Ctx) error {

	return utils.SendSuccessResponse(c, security.IpManager)
}

func RemoveBlockedIp(c *fiber.Ctx) error {

	return utils.SendSuccessResponse(c, "succesfully removed ip")
}
