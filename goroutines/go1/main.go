package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World")
}

func greeter(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(5 * time.Microsecond)
	}
}
