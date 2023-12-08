package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRankRoutes() {
	//Create
	app.Post("/rank/create", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Read
	app.Get("/rank/read", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Update
	app.Post("/rank/update", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Delete
	app.Delete("/rank/remove", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

}
