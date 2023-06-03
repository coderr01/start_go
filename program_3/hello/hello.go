package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	/*
		1. Set proerties of predefined logger.
		2. Set log entry prefix.
		3. Set flag to disable printing.
		4. Log includes
			a. Time
			b. Source File
			c. Line Number
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("Khushal")
	/*
		If error returned print it to console and exit
	*/
	if err != nil {
		log.Fatal(err)
	}
	// In case of no error return message to console
	fmt.Println(message)
}
