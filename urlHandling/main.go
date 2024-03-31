package main

import (
	"fmt"
	"net/url"
)

// Define a constant URL string
const myURL string = "https://go.dev:4000/solutions/google?feature=concurrency&loginid=99ff36"

func main() {
	fmt.Println("URL handling in Golang")
	fmt.Println(myURL)

	// Parse the URL
	result, err := url.Parse(myURL)
	Error(err)

	// Print various components of the parsed URL
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Port():", result.Port())

	// Get query parameters as a map
	queryParameters := result.Query()
	fmt.Printf("The type of query params are: %T\n", queryParameters) // It's a map of string slices (key-value pairs)
	fmt.Println(queryParameters["feature"])                           // Access specific query parameters
	fmt.Println(queryParameters["loginid"])

	// Iterate over query parameters
	for key, val := range queryParameters {
		fmt.Println(key, "->", val)
	}

	// Generate a URL from parts
	composeURL()
}

// Handle errors
func Error(err error) {
	if err != nil {
		panic(err)
	}
}

// Compose a URL from parts
func composeURL() {
	// Create URL parts struct
	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "go.dev",
		Path:    "/solutions/google",
		RawPath: "user=shrey",
	}
	// Convert URL parts to a string
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
