package loginhandlers

import (
	"api/storage"
	"api/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	//Generating a unique session identifier for further authentication
	sessionID := utils.NewUniqueIDGenerator().GenerateID()
	UserID := len(storage.SessionManager) + 1

	//Defining all userData
	userData := storage.UserData{UserID: UserID, Token: sessionID}

	// Store session data in the session manager
	storage.SessionManager[UserID] = userData

	// Return response with session ID in header or cookie
	returnValue := fmt.Sprintf("Login successful. Information: sessionID: %d, UserID: %d", sessionID, UserID)

	return c.SendString(returnValue)
}

func Logout(c *fiber.Ctx) error {

	user, err := utils.ParseRequestBody(storage.UserData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	// user := new(storage.UserData)

	// //Parsing the userData struct into the bodyParser to get the inserted values
	// if err := c.BodyParser(user); err != nil {
	// 	fmt.Println("error = ", err)
	// 	return c.SendStatus(200)
	// }

	//Check if the user is in the session manager
	// _, ok := storage.SessionManager[user.UserID]

	// if ok {
	// 	delete(storage.SessionManager, user.UserID)
	// 	return c.SendString("Successfully logged out")
	// }

	return c.Status(404).SendString("Failed to log out, user is not logged in")
}

func GetUsers(c *fiber.Ctx) error {

	bs, _ := json.Marshal(storage.SessionManager)
	fmt.Println(string(bs))

	return c.SendString("Printed all users")
}
