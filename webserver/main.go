package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SimpleServer)
	http.ListenAndServe(":8080", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	// Check if the length of the URL path is greater than 1
	if len(r.URL.Path) > 1 {
		// If the URL path is non-empty, respond with "Hello, Golang!" followed by the URL path starting from the second character
		fmt.Fprintf(w, "Hello, Golang! %s", r.URL.Path[1:])
	} else {
		// If the URL path is empty (i.e., root path), respond with "Hello, GoLang!"
		fmt.Fprintf(w, "Hello, GoLang!")
	}
}
