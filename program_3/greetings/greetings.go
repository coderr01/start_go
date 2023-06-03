package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greetings for the named person
func Hello(name string) (string, error) {
	// If no name was given, return error message
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create  a message using random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

/*
	randomFormat returns one of the set of greeting message.
	The returned message is selected randomly
*/

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"It's nice to meet you, %v!",
		"Hello %v! we meet again",
	}
	randomNum := rand.Intn(len(formats))
	return formats[randomNum]
}
