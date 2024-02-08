package routes

import (
	loginhandlers "api/handlers/login"
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

	// app.Use(func(c *fiber.Ctx) {

	// 	header := c.Request().Header
	// 	key := os.Getenv("SECRET_KEY")

	// 	if header != key {
	// 		return err
	// 	}

	// 	c.Next()
	// })

	//Login
	app.Get("/login", loginhandlers.Login)
	app.Get("/logout", loginhandlers.Logout)
	app.Get("/sessionStorage", loginhandlers.GetUsers)

	user := app.Group("/user", middleware.SessionAuth)
	user.Get("/get", userhandlers.Get)

	// rank := app.Group("/rank", authentication.AuthMiddleware)
	// rank.Get("/get", rankhandlers.Get)
}
