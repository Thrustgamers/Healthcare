package utils

import (
	"errors"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Body interface{}
}

func ParseRequestBody(class interface{}, c *fiber.Ctx) (Response, error) {

	if class == nil {
		return Response{}, errors.New("empty class given")
	}

	// Creating an instance based on the type of the provided class
	ResponseData := reflect.New(reflect.TypeOf(class)).Interface()

	// Parsing the userData struct into the bodyParser to get the inserted values
	if err := c.BodyParser(ResponseData); err != nil {
		return Response{}, err
	}

	response := Response{Body: ResponseData}
	return response, nil
}
