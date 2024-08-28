package loginhandlers

import (
	db "api/database"
	"api/models"
	"api/storage"
	"api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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

	//locally define DB variable
	db := db.Database.Db

	// Check Database for correct credentials and password
	// Also check for rank of person
	dbData := &models.Users{}
	err := db.Preload("RankRefer").Where(&models.Users{EmployeeId: user.EmployeeId, Password: user.Password}).First(dbData).Error

	if err != nil {
		log.Error().Err(err).Msg("Invalid credentials")
		return utils.SendErrorResponse(c, 401, errors.New("invalid credentials"))
	}

	// Check if session already exists
	if storage.DoesSessionExist(dbData.EmployeeId) {
		log.Warn().Msg("Login attempted on already active session")
		return utils.SendErrorResponse(c, 404, errors.New("error occured, session already active (contact support)"))
	}

	// Generating a unique session identifier for further authentication
	sessionID := uuid.New()
	UserID := dbData.ID

	// Defining all userData
	userData := storage.UserSession{
		UserID:     UserID,
		Token:      sessionID,
		Admin:      dbData.RankRefer.Admin == "YES",
		EmployeeId: dbData.EmployeeId,
	}

	// Serialize user data to JSON
	userDataJSON, err := json.Marshal(storage.UserSession{
		UserID: UserID,
		Token:  sessionID,
	})
	if err != nil {
		log.Fatal().Msg("Serializing user data to json errored")
	}

	// Store session data in the session manager
	storage.SessionManager[UserID] = userData

	// Creating authentication cookie
	cookie := &fiber.Cookie{
		Name:    "authentication",
		Value:   string(userDataJSON),
		Expires: time.Now().Add(24 * time.Hour),
		Secure:  true,
	}

	log.Info().Msg(fmt.Sprintf("Login successful. Information: sessionID: %d, UserID: %d", sessionID, UserID))

	return utils.SendSuccessResponse(c, loginDetail{
		Name:       dbData.Name,
		EmployeeId: dbData.EmployeeId,
	}, cookie)
}

func GetUsers(c *fiber.Ctx) error {
	bs, _ := json.Marshal(storage.SessionManager)
	fmt.Println(string(bs))

	return c.SendString("Printed all users")
}
