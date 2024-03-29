package main

import "fmt"

func main() {
	// Print a welcome message
	fmt.Println("Welcome to the language class!")

	// Create an empty map named 'language' with string keys and string values
	language := make(map[string]string)

	// Add entries to the 'language' map
	language["Go"] = "Golang"
	language["Py"] = "Python"
	language["Kt"] = "Kotlin"
	language["Js"] = "Javascript"

	// Print the entire 'language' map
	fmt.Println(language)

	// Print the full name associated with the key "Go"
	fmt.Println("Go is short for: ", language["Go"])

	// Print the type of the 'language' map
	fmt.Printf("The map is of the type %T:\n ", language)

	//use delete keyword to delete the values in maps
	delete(language, "Py")
	fmt.Println("After deleting:", language)

	//loops in golang
	for key, value := range language {
		fmt.Printf("The key is %v and value is %v!\n", key, value)
	}

}
