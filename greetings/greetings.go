package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		// go can return multiple values
		return "", errors.New("empty name")
	}

	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	// := -> declaring & initilizing a var in one line
	// %v -> format verb
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	//return string
	return message, nil
}
