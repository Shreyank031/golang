package main

import (
	"net/http"
	"sync"

	"fmt"
)

var wg sync.WaitGroup //Wait.Group is the modified/enhanced/advanced version of time.Sleep()
//These are pointers

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
}

func getStatus(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops...Not reachable!")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
