package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes() {
	//Create
	app.Post("/user/create", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Read
	app.Get("/user/read", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Update
	app.Post("/user/update", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Delete
	app.Delete("/user/remove", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

}
