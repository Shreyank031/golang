# Sending POST Request with Form Data in Go

This Go program demonstrates how to send a POST request with form data using the standard `net/http` package.

## Code Explanation

### 1. Define the URL Constant and Main Function

``` go
const Url string = "http://localhost:8000/postform"

func main() {
	fmt.Println("Post form data in golang")
	postFormRequest(Url)
}
```

- The main function is the entry point of the program.
- It prints a message to indicate the intention to perform a POST request.
- Then, it calls the postFormRequest function with the specified URL.

### 2. Define the `postFormRequest` Function

```go
func postFormRequest(Url string) {
	// Create form data using url.Values
	data := url.Values{}
	data.Add("First_name", "shreyank")
	data.Add("Last_name", "cm")
	data.Add("Mail", "shrey@go.dev")
```

- The postFormRequest function is responsible for sending a POST request with form data.
- It creates form data using `url.Values`.
- Three key-value pairs are added to the form data representing first name, last name, and email.

### 3. Send POST request with form data

```go
	responce, err := http.PostForm(Url, data)
	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()
 ```
-  Sends a POST request to the specified URL using `http.PostForm`.
-  Provides the URL and form data as arguments.
-  Performs error handling, panics if an error occurs.
-  Ensures the response body is closed after function execution using defer.

### 4. Read response, write, and print responce content.

```go
	content, err := io.ReadAll(responce.Body)
	var str strings.Builder
	dataByte, err := str.Write(content)
	if err != nil {
		panic(err)
	}
	// Print response content
	fmt.Println("Number of dataBytes:", dataByte)
	fmt.Println(str.String())
}
```

## Conclusion
This breakdown explains each part of the code and its functionality, providing a clear understanding of how the program sends a POST request with form data in Go.
