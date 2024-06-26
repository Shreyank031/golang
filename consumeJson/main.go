package main

import (
	"encoding/json"
	"fmt"
)

// course struct represents a course with its details
type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// DecodeJson decodes JSON data representing a course into a Go struct and map
func DecodeJson() {
	// JSON data to decode
	jsonDataFromWeb := []byte(`
		{
			"coursename": "DevOps",
			"price": 2999,
			"website": "youtube.com",
			"tags": ["kubernetes","Docker"]
		}
	`)

	var onlineCource course // Declare a variable to hold decoded course data

	// Check if the JSON data is valid
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")

		// Unmarshal the JSON data into the onlineCourse struct
		err := json.Unmarshal(jsonDataFromWeb, &onlineCource)
		if err != nil {
			panic(err)
		}

		// Print the decoded course struct
		fmt.Printf("%#v\n", onlineCource)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// Decode the JSON data into a map[string]interface{}
	//We know that key is strings but value can be any type,so we use interface-> it can hold any type
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)

	// Print the decoded map and its key-value pairs
	fmt.Printf("\n%#v\n", onlineData)
	for key, value := range onlineData {
		fmt.Printf("\nThe key is %v and value is %v and Type %T: \n", key, value, value)
	}
}

func main() {
	DecodeJson()
}
