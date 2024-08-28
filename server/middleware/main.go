package middleware

import (
	"api/storage"
	"api/utils"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type authentication struct {
	UserID uint      `json:"UserID,string"`
	Token  uuid.UUID `json:"Token,string"`
}

func SessionAuth(c *fiber.Ctx) error {
	// Retrieve session ID from cookies
	data := c.Cookies("authentication")

	//check if data is not a empty string
	if data != "" {

		var jsonData authentication

		err := json.Unmarshal([]byte(data), &jsonData)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if data, ok := storage.SessionManager[jsonData.UserID]; ok && data.Token == jsonData.Token {
			fmt.Printf("Authentication Succeeded on path %s", c.OriginalURL())
			return c.Next()
		} else {
			fmt.Println(ok)
			log.Warn().Msg("unauthorized login attempted")
			return utils.SendErrorResponse(c, 404, errors.New("authentication failed"))
		}

	}

	log.Warn().Msg("unauthorized request attempted")

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
			log.Warn().Msg("unauthorized login attempted")
			return utils.SendErrorResponse(c, 404, errors.New("authentication failed"))
		}

	}

	log.Warn().Msg("unauthorized admin request attempted")

	// Return unauthorized status if session is not found
	err := errors.New("admin authentication failed")
	return utils.SendErrorResponse(c, 404, err)
}
