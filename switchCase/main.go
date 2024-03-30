package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Cases in Golang")

	// Create a new random number generator seeded with the current time
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 6 (inclusive) representing the result of rolling a dice
	diceNumber := randGenerator.Intn(6) + 1

	fmt.Println("Value of dice is:", diceNumber)

	// Use a switch statement to perform different actions based on the value of the dice roll
	switch diceNumber {
	case 1:
		fmt.Println("Device value is one and you can open")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
	case 4:
		fmt.Println("You can move 4 steps")
		// After moving 4 steps, continue to the next case using fallthrough
		fallthrough
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps and roll the dice again")
	default:
		// If the dice roll is not between 1 and 6, print a default message
		fmt.Println("What was that?")
	}
}
