package adminRoutes

import (
	adminhandlers "api/handlers/admin"
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.AdminAuth)
	admin.Get("/", adminhandlers.GetBlockedIps)
}
