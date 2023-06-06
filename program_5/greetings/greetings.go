package greetings

import (
	"errors"
	"fmt"

	"math/rand"
)

// Hello functions return message from names person

func Hello(name string) (string, error) {
	// Error in case of No name is given
  
	if name == "" {
		return name, errors.New("No name is given")
	}
	// Create message
	message := fmt.Sprintf(randomFormat(), name)
	//fmt.Printf(message + "\n")
	return message, nil
}

// Function `Hellos` return a map that associate the person with message

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with message
	messages := make(map[string]string)
	/*
		1. Loop through list of names received
		2. Call hello function to get message for each name
	*/
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
    }
		// Associate the message with name
		messages[name] = message
	}
	return messages, nil

}

// Function `randomFormat` returns a set of greeting message randomly

func randomFormat() string {
  // A slice of message formats.
	formats := []string{
		"Hi, %v. Welocme!",
		"Good to see you %v!",
		"Hey, %v!, we met again",
	}
	// Return one of the message formats selected at random
  return formats[rand.Intn(len(formats))]
}
