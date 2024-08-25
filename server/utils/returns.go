package utils

import "github.com/gofiber/fiber/v2"

type ServerResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func SendSuccessResponse(c *fiber.Ctx, data interface{}, cookies ...*fiber.Cookie) error {
	resp := ServerResponse{
		Status: "success",
		Data:   data,
	}

	for _, cookie := range cookies {
		c.Cookie(cookie)
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func SendErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(ServerResponse{
		Status: "error",
		Error:  err.Error(),
	})
}

func SendWrongCreditionals(c *fiber.Ctx, errorMessage string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(ServerResponse{
		Status: "wrong credentials",
		Error:  errorMessage,
	})
}
