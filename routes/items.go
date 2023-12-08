package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes() {
	//Create
	app.Post("/item/create", func(c *fiber.Ctx) error {
		fmt.Println("test")
		return c.SendString(string(c.Body()))
	})

	//Read all
	app.Get("/item/read", func(c *fiber.Ctx) error {
		fmt.Println("Call triggered")
		return c.SendString("test")
	})

	//Read specific
	app.Get("/item/read/:id", func(c *fiber.Ctx) error {
		return c.SendString("Test")
	})

	//Update
	app.Post("/item/update", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

	//Delete
	app.Delete("/item/remove", func(c *fiber.Ctx) error {
		fmt.Println(c.Body())
		return c.SendString(c.Params("id"))
	})

}
