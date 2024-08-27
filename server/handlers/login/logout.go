package loginhandlers

import (
	"api/storage"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {

	user := new(storage.UserSession)

	//Parsing the userData struct into the bodyParser to get the inserted values
	if err := c.BodyParser(user); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	// Check if the user is in the session manager
	_, ok := storage.SessionManager[user.UserID]

	// Remove user from the session manager
	if ok {
		delete(storage.SessionManager, user.UserID)
		return c.SendString("Successfully logged out")
	}

	return c.Status(404).SendString("Failed to log out, user is not logged in")
}
