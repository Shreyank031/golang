package main

import (
	"encoding/json"
	"fmt"
)

// course represents a course with its attributes.
type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Create Json in golang")
	EncodeJson()
}

// EncodeJson encodes course data into JSON format.
func EncodeJson() {
	// Define slice of course structs
	ytCourses := []course{
		{"DevOps", 2999, "youtube.com", "passMeVada", []string{"kubernetes", "Docker"}},
		{"Native_App_Dev", 1999, "youtube.com", "August22", []string{"jetpack_Components", "jetpack_Compose"}},
		{"Mern_Dev", 999, "youtube.com", "Ez", nil},
	}

	// Encode data into JSON format with proper indentation
	trueJson, err := json.MarshalIndent(ytCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", trueJson) // Print the JSON data
}
