package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
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

	//	return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
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
func createOneCourse(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Create one course")
	w.Header().Set("Content-type", "application/json")

	//what if the data is empty.
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	//what if the data is {}
	var course Course
	//decode the data getting from r.Body (responce)
	_ = json.NewDecoder(r.Body).Decode(course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON.")
		return
	}
	//generate an unique id for course id, convert to string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
