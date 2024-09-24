package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a struct to represent a person
type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a slice of possible names
	names := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}

	// Create a slice of possible genders
	genders := []string{"Male", "Female"}

	// Generate 10 random people
	people := make([]Person, 10)
	for i := range people {
		people[i] = Person{
			Name:   names[rand.Intn(len(names))],
			Age:    rand.Intn(100), // Age between 0 and 99
			Gender: genders[rand.Intn(len(genders))],
		}
	}

	// Print the generated people
	fmt.Println("Generated people:")
	for _, person := range people {
		fmt.Println(person)
	}
}
