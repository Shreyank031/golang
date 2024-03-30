package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	// Calling the 'add' function with two arguments
	result := add(3, 9)
	fmt.Println("The result is:", result)

	// Calling the 'ProAdd' function with a variable number of arguments
	// The returned values are captured in 'slicy' and 'strr'
	slicy, strr := ProAdd(5, 1, 9, 8)
	fmt.Println("Pro result is:", slicy)
	fmt.Println("Pro message is:", strr)
}

// Function 'add' takes two integers as arguments and returns their sum
func add(num1 int, num2 int) int {
	return num1 + num2
}

// function that can accept a varying number of arguments.
// Function 'ProAdd' takes a variadic number of integers as arguments
// It calculates the sum of all the integers and returns the sum along with a string message
func ProAdd(values ...int) (int, string) {
	total := 0
	// Iterating over the provided values and calculating the total sum
	for _, val := range values {
		total += val
	}
	str := "Pro addition!"
	return total, str
}
