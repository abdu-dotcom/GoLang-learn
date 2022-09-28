package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a greeting that embeds(menyematkan) the name in a message.
	// := merupakan deklarsi dan inisialisasi varibale dalam satu baru
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}