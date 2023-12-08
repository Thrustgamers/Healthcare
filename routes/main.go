package routes

import (
	"api/database"

	"github.com/gofiber/fiber/v2"
)

var DB = database.Database
var app *fiber.App

func SetupRoutes(data *fiber.App) {
	app = data
	SetupItemRoutes()
	SetupRankRoutes()
	SetupUserRoutes()
}
