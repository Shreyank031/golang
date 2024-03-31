package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// URL to make the GET request
const URL string = "http://localhost:8000/get"

// main function initiates the program
func main() {
	fmt.Println("Web verb ")
	// Perform a GET request to the specified URL
	PerformGetRequest(URL)
}

// PerformGetRequest makes a GET request to the given URL
func PerformGetRequest(url string) {
	// Make a GET request to the URL
	responce, err := http.Get(url)
	Error(err)                  // Handle the error if any
	defer responce.Body.Close() // Close the response body after the function returns

	// Print the status code and content length of the response
	fmt.Println("Status code:", responce.StatusCode)
	fmt.Println("Content length:", responce.ContentLength)

	// Initialize a strings.Builder to construct the response string
	var responceString strings.Builder

	// Read the response body into a byte slice
	content, err := io.ReadAll(responce.Body)
	Error(err)

	// Write the content of the response body to responceString using its Write method
	byteCount, err := responceString.Write(content)
	Error(err)
	fmt.Println("ByteCount is:", byteCount)

	// Print the content of the response body as a string
	fmt.Println("1", string(content))

	// Print the content of responceString
	fmt.Println("2", responceString.String())
}

// Error function checks if the error is non-nil and panics if so
func Error(err error) {
	if err != nil {
		panic(err)
	}
}
