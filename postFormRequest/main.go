package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const Url string = "http://localhost:8000/postform"

func main() {
	fmt.Println("Post form data in golang")
	postFormRequest(Url)
}

// postFormRequest sends a POST request with form data to the specified URL.
func postFormRequest(Url string) {
	// Create form data using url.Values.
	data := url.Values{}
	data.Add("First_name", "shreyank")
	data.Add("Last_name", "cm")
	data.Add("Mail", "shrey@go.dev")

	// Send a POST request with the form data.
	//Create a form data. The data can be accessed by "url" package. Initally we'll be creating empty
	//key value pair for that, ie. Value{}
	response, err := http.PostForm(Url, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // Close the response body when done.

	// Read the response body to retrieve the content.
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Create a string builder to write the response content.
	var str strings.Builder
	// Write the response content to the string builder.
	dataBytes, err := str.Write(content)
	if err != nil {
		panic(err)
	}

	// Print the number of bytes written and the response content.
	fmt.Println("Number of dataBytes:", dataBytes)
	fmt.Println(str.String())
}
