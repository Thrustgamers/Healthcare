package routes

import (
	adminhandlers "api/handlers/admin"
	loginhandlers "api/handlers/login"
	testhandlers "api/handlers/test"
	userhandlers "api/handlers/users"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func SetupRoutes(app *fiber.App) {

	//Preparing dotenv files
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err)
	}

	//Login
	app.Get("/login", loginhandlers.Login)
	app.Get("/logout", loginhandlers.Logout)
	app.Post("/statuscheck", loginhandlers.StatusCheck)

	//Users
	user := app.Group("/user", middleware.SessionAuth)
	user.Get("/get", userhandlers.Get)

	//Admin
	admin := app.Group("/admin", middleware.AdminAuth)
	admin.Get("/", adminhandlers.GetBlockedIps)

	//Test
	test := app.Group("/test")
	test.Get("/sessions", testhandlers.GetSessions)

	// rank := app.Group("/rank", authentication.AuthMiddleware)
	// rank.Get("/get", rankhandlers.Get)
}
