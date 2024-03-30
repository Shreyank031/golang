package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This file needs to go inside a file -> GogoGolang"
	file, err := os.Create("./createdFile.txt")
	Error(err) // Check for and handle errors
	// Write content to the file

	length, err := io.WriteString(file, content)
	Error(err) // Check for and handle errors
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./createdFile.txt")
}

// Error function handles errors
func Error(err error) {
	if err != nil {
		panic(err)
	}
}

// Read the file
func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	Error(err) // Check for and handle errors
	fmt.Println("The data inside the file is:\n", string(dataByte))
}
