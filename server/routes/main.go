package routes

import (
	loginhandlers "api/handlers/login"
	userhandlers "api/handlers/users"
	"api/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func SetupRoutes(app *fiber.App) {

	//Preparing dotenv files
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err)
	}

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println(fmt.Printf("Incoming request to host:%s", c.OriginalURL()))

		return c.Next()
	})

	//Login
	app.Get("/login", loginhandlers.Login)
	app.Get("/logout", loginhandlers.Logout)
	app.Post("/statuscheck", loginhandlers.StatusCheck)
	app.Get("/sessionStorage", loginhandlers.GetUsers)

	user := app.Group("/user", middleware.SessionAuth)
	user.Get("/get", userhandlers.Get)

	// rank := app.Group("/rank", authentication.AuthMiddleware)
	// rank.Get("/get", rankhandlers.Get)
}
