package rankRoutes

import (
	rankhandlers "api/handlers/ranks"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRankRoutes(app *fiber.App) {
	rank := app.Group("/rank", middleware.AdminAuth)
	rank.Get("/get", rankhandlers.Get)
}
