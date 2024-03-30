package main

import "fmt"

// Rectange represents a rectangle with length and width.
type Rectange struct {
	length, width float64
}

// User represents a user with name, email, and age.
type User struct {
	Name, Email string
	Age         int
}

/* Methods in Go are functions associated with a specific type, and they are declared with
a receiver parameter. The receiver parameter (r Rectangle) indicates that this area() method
belongs to the Rectangle type, and it can be called on instances of Rectangle. */

/* Since the method is associated with the Rectangle type, it can access the fields of the
Rectangle instance r directly. */

// area calculates the area of the rectangle.
func (r Rectange) area() float64 {
	return r.length * r.width
}

// name prints the email of the user.
func (u User) name() {
	fmt.Println("Email of this user is:", u.Email)
}

// newUser modifies the email of the user and returns the modified email.
func (m User) newUser() string {
	m.Email = "golang@go.dev"
	return m.Email
}

func main() {
	// Create a rectangle instance with specified length and width.
	value := Rectange{22.7, 11.9}
	// Print the area of the rectangle.
	fmt.Println(value.area())

	// Create a user instance with specified name, email, and age.
	shrey := User{"shrey", "shrey@go.dev", 22}
	// Print the email of the user.
	shrey.name()
	// Modify the email of the user and print the modified email.
	fmt.Printf("The modified email is: %v", shrey.newUser())
}
