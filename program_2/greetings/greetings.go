package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person 
func Hello (name string) (string, error) {
	// If input variable name is not given then return an error with message
	if name == "" {
		return "", errors.New("Input is not given")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
