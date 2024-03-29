package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	presentTime := time.Now()

	// Print the current time
	fmt.Println(presentTime)

	// Format and print the current time in a specific layout
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Create a specific date and time
	createDate := time.Date(2002, time.May, 31, 01, 15, 0, 0, time.UTC)

	// Print the created date and time
	fmt.Println(createDate)

	// Format and print the created date in a specific layout
	fmt.Println(createDate.Format("02-01-2006 Monday"))
}
