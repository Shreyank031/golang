package main

import "fmt"

func main() {
	// This will be executed first when the function returns
	defer fmt.Println("Hello")

	// This will be executed second when the function returns
	defer fmt.Println("One")

	// This will be executed third when the function returns
	defer fmt.Println("\nTwo")

	// This will be executed immediately
	fmt.Println("World!")

	// Call the myDefer function
	myDefer()
}

func myDefer() {
	for i := 0; i <= 6; i++ {
		// This will be executed in reverse order when the function returns
		defer fmt.Print(i)
	}
}
