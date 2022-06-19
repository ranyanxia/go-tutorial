package main

import (
	"fmt"
	"log"

	"roystudio.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	//if an error was returned, print it to console and exit the program
	if err != nil {
		//log.Fatal => print and exit
		log.Fatal(err)
	}

	// if no error, print the message
	fmt.Println(message)
}
