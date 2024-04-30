package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
	//"time"
)

func main() {

	// rand.Seed(time.Now().UnixMicro())
	// for i := 1; i <= 10; i++ {
	//
	// 	fmt.Printf("%d ", rand.Intn(21))
	// }
	// fmt.Println()
	// fmt.Println(rand.Intn(5))

	for i := 0; i <= 10; i++ {
		myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(55))
		fmt.Printf("%d ", myRandomNum)
	}
	fmt.Println()
}
