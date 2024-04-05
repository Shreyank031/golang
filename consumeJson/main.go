package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func DecodeJson() {

	jsonDataFromWeb := []byte(`

 {
                "coursename": "DevOps",
                "price": 2999,
                "website": "youtube.com",
                "tags": ["kubernetes","Docker"]
        }
	`)

	var onlineCource course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &onlineCource)
		fmt.Printf("%#v", onlineCource)
	} else {
		fmt.Println("JSON is not valid")
	}

}
func main() {
	DecodeJson()
}
