package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a function to generate a random password of a given length
func generatePassword(length int) string {
	// Define the characters to choose from
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()")

	// Create a slice to store the password
	password := make([]rune, length)

	// Generate random characters and add them to the password
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	// Return the password as a string
	return string(password)
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a password of length 16
	password := generatePassword(16)
	fmt.Println("Generated password:", password)
}
