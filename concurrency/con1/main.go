package main

import "fmt"

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()
	msg := <-myChannel

	fmt.Println(msg)
}
