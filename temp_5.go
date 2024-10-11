package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time.
	currentTime := time.Now()

	// Format the time.
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	// Print the formatted time.
	fmt.Println("Current time:", formattedTime)
}
