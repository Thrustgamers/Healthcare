package main

import (
	"api/database"
	"api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/highlight/highlight/sdk/highlight-go"
	highlightFiber "github.com/highlight/highlight/sdk/highlight-go/middleware/fiber"
)

func main() {
	database.ConnectToDb()

	highlight.SetProjectID("lgx91qqg")
	highlight.Start(highlight.WithServiceName("go_webshop"), highlight.WithServiceVersion("git-sha"))

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       "Healthcare Server v0.0.1",
	})

	app.Use(cors.New())
	app.Get("/metrics", monitor.New())
	// app.Static("/", "./web/dist/")
	app.Use(highlightFiber.Middleware())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
