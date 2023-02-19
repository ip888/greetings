package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) string {
	// If no name was given, return an error with message
	if name == "" {
		errors.New("Empty name")
	}
	// If a name was received, return a value that embeds the name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
