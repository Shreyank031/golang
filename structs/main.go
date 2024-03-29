package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	fmt.Println("No inheritance in Go. A bummer.")

	// Create a struct instance named 'shrey' with values assigned to its fields
	shrey := SuperUser{"Shreyank", 22, "shrey@go.dev", true}
	fmt.Println(shrey)
	// Print the 'shrey' struct instance with field names
	fmt.Printf("Full details -> %+v\n", shrey)
	// Print specific fields of the 'shrey' struct instance
	fmt.Printf("Name is %v and Email is %v", shrey.Name, shrey.Email)
}

// Define a struct named 'SuperUser' to represent a super user
// The name of the struct should be in UpperCase, i.e., 'SuperUser'
type SuperUser struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
