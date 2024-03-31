package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Define the URL to make a GET request to
	const URL = "https://github.com/Shreyank031"
	fmt.Println("Web Request")

	// Make a GET request to the URL
	request, err := http.Get(URL)
	if err != nil {
		panic(err) // Panic if there's an error making the request
	}
	fmt.Printf("The request is of the type: %T\n", request)

	// Close the response body to prevent resource leaks
	defer request.Body.Close() // Deferred close to ensure the body is closed after the function returns

	// Read the response body
	databytes, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	// Convert response body bytes to a string and print it
	content := string(databytes)
	fmt.Println(content)
}
