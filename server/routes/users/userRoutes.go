package userRoutes

import (
	userhandlers "api/handlers/users"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user", middleware.SessionAuth)
	user.Get("/", userhandlers.Get)
}
