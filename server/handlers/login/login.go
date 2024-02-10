package loginhandlers

import (
	"api/database"
	"api/models"
	"api/storage"
	"api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
)

type loginDetail struct {
	EmployeeId int    `json:"employeeId"`
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
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
	data := database.Database.Db.Where(&models.Users{EmployeeId: user.EmployeeId, Password: user.Password}).First(&models.Users{})

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
	returnData := loginDetail{
		Name: "test",
	}

	// Serialize user data to JSON
	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		log.Fatal().Msg("Serializing user data to json errored")
	}

	//Creating authentication cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "authentication"
	cookie.Value = string(userDataJSON)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = true

	log.Info().Msg(fmt.Sprintf("Login successful. Information: sessionID: %d, UserID: %d", sessionID, UserID))

	return utils.SendSuccessResponse(c, returnData, cookie)
}

func GetUsers(c *fiber.Ctx) error {

	bs, _ := json.Marshal(storage.SessionManager)
	fmt.Println(string(bs))

	return c.SendString("Printed all users")
}
