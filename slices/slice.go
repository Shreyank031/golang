package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declare and initialize a slice of strings
	var fruitList = []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("Type of fruitList is: %T\n", fruitList) // Print the type of fruitList
	fmt.Println(fruitList)                              // Print the elements of fruitList

	// Append new elements to fruitList
	fruitList = append(fruitList, "Berries", "Orange")
	fmt.Println(fruitList)

	// Slice fruitList to include elements from index 1 to 3 (exclusive) and reassign it to fruitList
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Declare and initialize a slice of integers with a length of 4
	highScores := make([]int, 4)

	// Assign values to elements of highScores
	highScores[0] = 901
	highScores[1] = 581
	highScores[2] = 195
	highScores[3] = 321

	// Append new elements to highScores
	highScores = append(highScores, 1099, 404, 433)
	fmt.Println(highScores)

	// Check if highScores is sorted
	fmt.Println(sort.IntsAreSorted(highScores))

	// Sort highScores
	sort.Ints(highScores)
	fmt.Println(highScores)

	// Check if highScores is sorted after sorting
	fmt.Println(sort.IntsAreSorted(highScores))
	// Slice the language slice to remove the element at index 2
	var language = []string{"golang", "rust", "python", "kotlin"}
	var index = 2
	fmt.Println(append(language[:index], language[index+1:]...))

}
