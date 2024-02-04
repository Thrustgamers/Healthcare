package main

import (
	"api/database"
	"api/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/rs/zerolog/log"
)

func main() {

	//Preparing database
	database.ConnectToDb()

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

	routes.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Error().Err(err)
		os.Exit(3)
	}

	log.Info().Msg("Server launched on port 3000")
}
