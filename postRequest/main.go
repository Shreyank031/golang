package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// URL represents the URL endpoint for the POST request.
const URL string = "http://localhost:8000/post"

func main() {
	fmt.Println("Post request in Golang")

	// Perform a POST request with JSON data to the specified URL.
	PerformPostJsonRequest(URL)
}

// PerformPostJsonRequest sends a POST request with JSON data to the specified URL.
func PerformPostJsonRequest(url string) {
	// Create the data to be sent in the request body.

	//first create the data you want to send. In our case we'll be creating some payload dummy data
	//fake payload json

	//You can crate any kind data in strings.NewReader(``)

	requestBody := strings.NewReader(`
		{
		"programmingLanguage":"Golang",
		"inShort":"Go",
		"we_call_them":"gophers"
		}
	`)

	// Send the POST request with the JSON data.
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // Close the response body when done.

	// Read the response body to retrieve the content.
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Print the content of the response body.
	fmt.Println(string(content))
}
