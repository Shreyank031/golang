package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("All about Channels in golang")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 9
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
