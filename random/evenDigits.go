package main

import "fmt"

func main() {

	array := []int{23, 87, 3432, 1, 872, 986753, 380}
	fmt.Println(array)
	fmt.Println(findNumbers(array))
}
func findNumbers(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if even(arr[i]) {
			res++
		}
	}
	return res
}
func even(nums int) bool {
	x := 0
	for nums > 0 {
		nums = nums / 10
		x++
	}
	return x%2 == 0
}
