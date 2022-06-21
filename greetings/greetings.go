package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// Loop the slice of names, and call the Hello function to get msg for each name
	//for range, _ is 'blank identifier'
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// Keep this one to make the module backward compatibility
func Hello(name string) (string, error) {
	if name == "" {
		// go can return multiple values
		return "", errors.New("empty name")
	}

	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	// := -> declaring & initilizing a var in one line
	// %v -> format verb
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat(), name)

	//return string
	return message, nil
}

// init -> Go execute init automatically at program startup, after global vars have been initialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

// func with lower case is not accessible by outside of the package
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

// try maps: https://go.dev/blog/maps
func TryMap() map[string]int {
	//define map
	// var m map[string]int
	m := make(map[string]int)

	m["route"] = 66
	i := m["route"]
	fmt.Println(i)

	//set zero if not exist
	// j == 0
	j := m["root"]
	fmt.Println(j)

	// return length
	n := len(m)
	fmt.Println(n)

	// delete to remove element, return nothing. and do nothing if key not exist
	delete(m, "route")

	// assign value of "route" to i if exists, otherwise 0; and ok == true if exists, false if not
	i, ok := m["route"]
	fmt.Println(i, ok)

	//test a key without return value
	_, yes := m["route"]
	fmt.Printf("%v", yes)

	//iterate
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// init map with values
	commits := map[string]int{
		"a": 12,
		"b": 22,
		"c": 113,
	}
	fmt.Printf("%v", commits)

	return m
}
