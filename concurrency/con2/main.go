package main

import "fmt"

func main() {
	myChannel := make(chan string)
	myAnotherChannel := make(chan string)

	go func() {
		myAnotherChannel <- "anotherData"
	}()

	go func() {
		myChannel <- "data"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-myAnotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}

}
