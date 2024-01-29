package routes

import (
	loginhandlers "api/handlers/login"
	userhandlers "api/handlers/users"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//Login
	app.Get("/login", loginhandlers.Login)
	app.Get("/logout", loginhandlers.Logout)
	app.Get("/sessionStorage", loginhandlers.GetUsers)

	user := app.Group("/user", middleware.SessionAuth)
	user.Get("/get", userhandlers.Get)

	// rank := app.Group("/rank", authentication.AuthMiddleware)
	// rank.Get("/get", rankhandlers.Get)
}
