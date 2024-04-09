package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")

	fmt.Printf("Please rate our app between 1 and 5: ")

	// Create a reader to read user input from stdin
	reader := bufio.NewReader(os.Stdin)

	// Read user input until the newline character ('\n')
	input, _ := reader.ReadString('\n')

	// Trim leading and trailing whitespace from the input
	input = strings.TrimSpace(input)

	fmt.Println("Thanks for rating:", input)

	// Parse the input string as a floating-point number (float64)
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		// If parsing fails, print the error
		fmt.Println("Error:", err)
	} else {
		// If parsing succeeds, add 1 to the rating and print the updated rating
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
