package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/highlight/highlight/sdk/highlight-go"
	highlightFiber "github.com/highlight/highlight/sdk/highlight-go/middleware/fiber"
)

var app *fiber.App

func main() {
	database.ConnectToDb()

	highlight.SetProjectID("lgx91qqg")
	highlight.Start(highlight.WithServiceName("go_webshop"), highlight.WithServiceVersion("git-sha"))

	app = fiber.New()
	app.Use(cors.New())
	app.Get("/metrics", monitor.New())
	app.Static("/", "./web/dist/")
	app.Use(highlightFiber.Middleware())

	routes.SetupRoutes(app)

	app.Listen(":80")
}
