package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /path", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have hit the path")
	})
	router.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "The task with id: %s", id)
	})
	log.Fatal(http.ListenAndServe("localhost:9091", router))
}
