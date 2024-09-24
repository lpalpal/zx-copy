package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between    1 and 100
	randomNumber := rand.Intn(100) + 1
	fmt.Println("Random number:", randomNumber)   

	// Generate a random float between 0 and 1
	randomFloat := rand.Float64()
	fmt.Println("Random float:", randomFloat)

	// Generate a random string of length 10
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, 10)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]   
	}
	fmt.Println("Random string:", string(randomString))

	// Generate a random boolean
	randomBool := rand.Intn(2) == 0
	fmt.Println("Random boolean:", randomBool)
}
