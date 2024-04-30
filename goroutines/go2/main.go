package main

import (
	"net/http"
	"sync"

	"fmt"
)

// These are pointers
var signals = []string{"test"}
var wg sync.WaitGroup //Wait.Group is the modified/enhanced/advanced version of time.Sleep(). pointers
var mut sync.Mutex    //pointer

func main() {

	websiteList := []string{
		"https://github.com/Shreyank031",
		"https://google.com",
		"https://in.pinterest.com",
		"https://archlinux.com",
		"https://archlinux.org",
		"https://go.dev",
	}
	for _, web := range websiteList {
		go getStatus(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatus(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops...Not reachable!")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
