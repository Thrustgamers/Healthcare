package loginRoutes

import (
	loginhandlers "api/handlers/login"

	"github.com/gofiber/fiber/v2"
)

func SetupLoginRoutes(app *fiber.App) {
	app.Get("/login", loginhandlers.Login)
	app.Get("/logout", loginhandlers.Logout)
	app.Post("/statuscheck", loginhandlers.StatusCheck)
}
