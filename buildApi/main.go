package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //pointer pointing towards custom struct Author.
}
type Author struct {
	AuthorName string `json:"fullname"`
	Website    string `json:"website"`
}

// fake db using slice
var courses []Course //slice of the type Course

func (c *Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Building api in golang")
}

// Controller - file
// serve route
func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Shrey</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	//add a header
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getIOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-type", "application/json")

	//grab id from request provided by user
	params := mux.Vars(r)

	//loop through the courses. find the matching id and return the responce
	for _, course := range courses {
		if course.CourseId == params["id"] { // "id" something to do with route
			json.NewEncoder(w).Encode(course) //send the whole course
			return

		}

	}
	json.NewEncoder(w).Encode("No Course found with given id")
}
