package loginhandlers

import (
	"api/database"
	"api/models"
	"api/storage"
	"api/utils"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type loginDetail struct {
	EmployeeId int
	Password   string
}

func Login(c *fiber.Ctx) error {

	user := new(loginDetail)

	//Parsing the loginDetail struct into the bodyParser to get the inserted values
	if err := c.BodyParser(user); err != nil {
		fmt.Println("error = ", err)
		return utils.SendErrorResponse(c, 404, err)
	}

	fmt.Println(user)

	//Check Database for correct credentials and password
	data := database.Database.Db.Where(&models.Users{EmployeeId: user.EmployeeId, Password: user.Password})

	fmt.Println(data.RowsAffected)

	//If no records send error response to frontend
	if data.RowsAffected == 0 {
		err := errors.New("invalid credentials")
		return utils.SendErrorResponse(c, 404, err)
	}

	//Generating a unique session identifier for further authentication
	sessionID := utils.GenerateUniqueID()
	UserID := len(storage.SessionManager) + 1

	//Defining all userData
	userData := storage.UserData{UserID: UserID, Token: sessionID}

	// Store session data in the session manager
	storage.SessionManager[UserID] = userData

	//Define the data that is getting send back to the frontend
	returnData := storage.UserData{
		UserID: UserID,
		Token:  sessionID,
	}

	fmt.Printf(fmt.Sprintf("Login successful. Information: sessionID: %d, UserID: %d", sessionID, UserID))

	return utils.SendSuccessResponse(c, returnData)
}

func GetUsers(c *fiber.Ctx) error {

	bs, _ := json.Marshal(storage.SessionManager)
	fmt.Println(string(bs))

	return c.SendString("Printed all users")
}
