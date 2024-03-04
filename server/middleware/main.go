package middleware

import (
	"api/storage"
	"api/utils"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type authentication struct {
	UserID int `json:"UserID,string"`
	Token  int `json:"Token,string"`
}

func SessionAuth(c *fiber.Ctx) error {
	// Retrieve session ID from cookies
	data := c.Cookies("authentication")

	if data != "" {

		var jsonData authentication

		err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if data, ok := storage.SessionManager[jsonData.UserID]; ok && data.Token == jsonData.Token {
			fmt.Println("Authentication Succeeded")
			c.Next()
			return utils.SendSuccessResponse(c, "")
		} else {
			fmt.Println(ok)
			log.Warn().Msg("Unautherized login attempted")
			return utils.SendErrorResponse(c, 404, errors.New("authentication failed"))
		}

	}

	log.Warn().Msg("Unautherized request attempted")

	// Return unauthorized status if session is not found
	return utils.SendErrorResponse(c, 404, errors.New("authentication failed"))
}

func AdminAuth(c *fiber.Ctx) error {
	// Retrieve session ID from cookies
	data := c.Cookies("authentication")

	if data != "" {

		var jsonData authentication

		err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if data, ok := storage.SessionManager[jsonData.UserID]; ok && data.Token == jsonData.Token && data.Admin {
			fmt.Println("Authentication Succeeded")
			c.Next()
			return utils.SendSuccessResponse(c, "")
		} else {
			fmt.Println(ok)
			log.Warn().Msg("Unautherized login attempted")
			return utils.SendErrorResponse(c, 404, errors.New("authentication failed"))
		}

	}

	log.Warn().Msg("Unautherized admin request attempted")

	// Return unauthorized status if session is not found
	err := errors.New("admin authentication failed")
	return utils.SendErrorResponse(c, 404, err)
}
