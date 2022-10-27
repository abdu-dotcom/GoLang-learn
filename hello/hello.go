package main

import (
	"fmt"

	"example.com/greetings" // for give access function in those packages
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}