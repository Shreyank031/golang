package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, Golang!")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Oopsy"))
			http.Error(w, "Oopsy", http.StatusBadRequest)
			return
		}
		//		log.Printf("Data -> %s\n", d)
		fmt.Fprintf(w, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye, Golang!")
	})
	http.ListenAndServe("localhost:9090", nil)
}
