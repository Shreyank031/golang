package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	// Create a slice of days of the week
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// Loop over the slice using a traditional for loop with index variable
	// for d := 0; d < len(days); d++ {
	//     fmt.Println(days[d])
	// }

	// Loop over the slice using a range-based for loop with index
	// for i := range days {
	//     fmt.Println(days[i])
	// }

	// Loop over the slice using a range-based for loop with index and value
	// for _, day := range days {
	//     fmt.Printf("The index is blank and the value is %v.\n", day)
	// }

	// Loop over the slice using a range-based for loop with index and value
	for ind, day := range days {
		fmt.Printf("The index is %v and the value is %v.\n", ind, day)
	}

	value := 1

	// Start a loop that continues until "value" is less than 10
	for value < 10 {
		// Check if "value" is 7
		if value == 7 {
			// If "value" is 7, jump to the label "GogoGolang"
			goto GogoGolang
		}

		// Check if "value" is 5
		if value == 5 {
			// If "value" is 5, print a message and skip to the next iteration
			fmt.Println("Iam not printing 5 by the way, that's what continue does, it skips")
			value++
			continue
			//break
		}

		// Print the value of "value"
		fmt.Printf("The value is: %v\n", value)

		// Increment "value" by 1
		value++
	}

	// Label "GogoGolang" for the goto statement
GogoGolang:
	// Print a message after the goto statement
	fmt.Println("Go-to statement baby!")
}
