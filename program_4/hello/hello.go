package main

import (
  "fmt"
	"log"

	"example.com/greetings"
)

func main() {
	/*
		1. Set the predefined properties of logger
			a. Log entry prefix
			b. flag to disable printing
			c. Time, source file, line number
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Joe", "Kira", "Kitti"}

	// Request greeting message of names
	messages, err := greetings.Hellos(names)

	// In case of error
	if err != nil {
		log.Fatal(err)
	}
	_ = messages
	fmt.Println(messages)
}
