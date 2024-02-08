package utils

import "github.com/gofiber/fiber/v2"

type ServerResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func SendSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(ServerResponse{
		Status: "success",
		Data:   data,
	})
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(ServerResponse{
		Status: "error",
		Error:  err.Error(),
	})
}
