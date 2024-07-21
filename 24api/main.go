package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

// Course model -> separate file
type Course struct {
	CourseId   string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

// Author model -> separate file
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// IsEmpty middleware -> separate file
func (course *Course) IsEmpty() bool {
	//return course.CourseId == "" && course.CourseName == ""
	return course.CourseName == ""
}

func main() {

}

// controllers -> separate file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to apis in Go<h1/>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(courses)
	if err != nil {
		panic(err)
	}
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseId := params["id"]
	for _, v := range courses {
		if v.CourseId == courseId {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	jsonResponse := fmt.Sprintf(`{"message": "No course found with given id", "course_id": "%s"}`, courseId)
	json.NewEncoder(w).Encode(jsonResponse)
	return
}

func createCourseRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses := append(courses, course)
	json.NewEncoder(w).Encode(courses)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil {
				panic(err)
			}

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func removeCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseId := params["id"]
	// find the course with the specific id
	for index, val := range courses {
		if val.CourseId == courseId {
			// remove the course from the courses
			courses = append(courses[:index], courses[index+1:]...)
			break
		} else {
			panic("Could not find a course with given Id")
		}
	}
}
