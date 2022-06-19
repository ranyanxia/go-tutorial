package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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
	message := fmt.Sprintf(randomFormat(), name)

	//return string
	return message, nil
}

// init -> Go execute init automatically at program startup, after global vars have been initialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats
	// slice is like an array, except the size changes dynamically as you add & remove items
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
