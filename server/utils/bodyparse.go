package utils

import (
	"fmt"
)

func ParseRequestBody(class interface{}) (string, error) {

	// classType := reflect.TypeOf(class)

	fmt.Println(class)

	// user := new(classType)

	// //Parsing the userData struct into the bodyParser to get the inserted values
	// if err := c.BodyParser(user); err != nil {
	// 	fmt.Println("error = ", err)
	// 	return nil
	// }

	return "test", nil
}
