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

	// Parsing the loginDetail struct into the bodyParser to get the inserted values
	if err := c.BodyParser(user); err != nil {
		log.Error().Err(err)
		return utils.SendErrorResponse(c, 404, err)
	}

	// Check Database for correct credentials and password
	dbData := &models.Users{}
	data := database.Database.Db.Where(&models.Users{EmployeeId: user.EmployeeId, Password: user.Password}).First(dbData)

	// If no records send error response to frontend
	if data.RowsAffected == 0 {
		err := errors.New("invalid credentials")
		return utils.SendErrorResponse(c, 404, err)
	}

	// Check if session already exists
	if storage.DoesSessionExist(dbData.EmployeeId) {
		log.Warn().Msg("Login attempted on already active session")
		return utils.SendErrorResponse(c, 404, errors.New("error occured, session already active (contact support)"))
	}

	// Generating a unique session identifier for further authentication
	sessionID := utils.GenerateUniqueID()
	UserID := len(storage.SessionManager)

	dbRankData := &models.Ranks{}
	database.Database.Db.Where(&models.Ranks{ID: uint(dbData.Rank)}).First(dbRankData)

	fmt.Println(dbRankData)

	isAdmin := dbRankData.Admin == "YES"

	// Defining all userData
	userData := storage.UserData{
		UserID:     UserID,
		Token:      sessionID,
		Admin:      isAdmin,
		EmployeeId: dbData.EmployeeId,
	}

	// Serialize user data to JSON
	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		log.Fatal().Msg("Serializing user data to json errored")
	}

	// Store session data in the session manager
	storage.SessionManager[UserID] = userData

	// Define the data that is getting send back to the frontend
	returnData := loginDetail{
		Name:       dbData.Name,
		EmployeeId: dbData.EmployeeId,
	}

	// Creating authentication cookie
	cookie := &fiber.Cookie{
		Name:    "authentication",
		Value:   string(userDataJSON),
		Expires: time.Now().Add(24 * time.Hour),
		Secure:  true,
	}

	log.Info().Msg(fmt.Sprintf("Login successful. Information: sessionID: %d, UserID: %d", sessionID, UserID))

	return utils.SendSuccessResponse(c, returnData, cookie)
}

func GetUsers(c *fiber.Ctx) error {
	bs, _ := json.Marshal(storage.SessionManager)
	fmt.Println(string(bs))

	return c.SendString("Printed all users")
}
