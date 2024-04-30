package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition -> go.dev")

	wg := &sync.WaitGroup{}

	//Whenever you try to access a memory, you need to lock it so that all the time when you are performing operations on that memory nobody can come and mess around with the job :) -> mutex lock
	mut := &sync.Mutex{} //here whenever there is write operations use it.

	// m := &sync.RWMutex{}
	//
	// m.RLock()
	var score = []int{0}
	//	m.RUnlock()

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("goroutine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("goroutine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("goroutine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

//immediately invoked functions(ifs) . There are Ananymous/lamda functions you are executing them

/*
In Go, a race condition occurs when two or more goroutines access shared data concurrently, and at least one of them modifies the data. In the provided code, a race condition can occur due to concurrent access and modification of the score slice by multiple goroutines.

Here's how it can lead to a race condition:

    - Three goroutines are launched concurrently using the go keyword, each of which appends a different value (1, 2, or 3) to the score slice.
    - Since the score slice is shared among these goroutines and there is no synchronization mechanism in place to control access to it, they may execute concurrently and access/modify the score slice simultaneously.
    - If two or more goroutines attempt to append elements to the score slice simultaneously, it can lead to unexpected behavior, such as overwriting data or inconsistent state of the slice.
*/
