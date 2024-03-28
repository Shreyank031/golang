package main

import "fmt"

func main() {

	fmt.Println("Welcome to my quize game")

	fmt.Printf("Enter your name: ")
	name := "Shrey"
	fmt.Scan(&name)
	fmt.Printf("Hello %v, Welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)
	if age >= 10 {
		fmt.Println("Yay, you can play the game!")
	} else {
		fmt.Println("You cannot play the game")
		return
	}
	fmt.Printf("Which is more Powerfull, RTX 4080 or RTX 4090? ")
	var answer1 string
	var answer2 string
	var count int
	no_of_questions := 2

	fmt.Scan(&answer1, &answer2)

	if answer1+" "+answer2 == "RTX 4090" || answer1+" "+answer2 == "rtx 4090" {
		fmt.Println("CORRECT!")
		count++
	} else {
		fmt.Println("INCORRECT!")
	}

	var cores int
	fmt.Println("How many cores does a Ryzen 5 4500u has? ")
	fmt.Scan(&cores)

	if cores == 6 {
		fmt.Println("Correct!!")
		count += 1
	} else {
		fmt.Println("Incorrect")
	}
	fmt.Printf("You have scored %v out of %v\n", count, no_of_questions)
	percent := (float64(count) / float64(no_of_questions) * 100)
	fmt.Printf("You scored %v%%.", percent)

}
