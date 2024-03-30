package main

import "fmt"

func main() {
	fmt.Println("If else statements in Golang")

	logins := 31

	// Use if-else statements to check the value of logins
	if logins > 22 {
		fmt.Println("SuperUser")
	} else if logins < 22 {
		fmt.Println("Normal User")
	} else {
		fmt.Println("Sussy user")
	}

	// Use if-else statements to check if a number is even or odd
	if 11%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	// Use if-else statements with initialization to check if a number is even or odd
	if num := 12; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
