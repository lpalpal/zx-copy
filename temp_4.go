package main

import (
        "fmt"
        "math/rand"
        "time"
)

func main() {
        // Seed the random number generator with the current time
        rand.Seed(time.Now().UnixNano())

        // Generate a random number between    1 and 100
        randomNumber := rand.Intn(100)    + 1

        // Print the random number
        fmt.Println("Your random number is:", randomNumber)
}
