package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("All about Channels in golang")

	// Create a buffered channel with a capacity of 2
	myCh := make(chan int, 2) //two channels
	// Create a WaitGroup to synchronize goroutines
	wg := &sync.WaitGroup{}

	wg.Add(2) // Increment the WaitGroup counter by 2 for the two goroutines
	//recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpenn := <-ch
		fmt.Println(isChannelOpenn)
		fmt.Println(val)
		v, c := <-ch
		fmt.Println(c)
		fmt.Println(v)

		wg.Done()
	}(myCh, wg)

	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 9
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
