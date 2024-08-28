package testRoutes

import (
	testhandlers "api/handlers/test"

	"github.com/gofiber/fiber/v2"
)

func SetupTestRoutes(app *fiber.App) {
	test := app.Group("/test")
	test.Get("/sessions", testhandlers.GetSessions)
}
