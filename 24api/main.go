package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
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
	r := mux.NewRouter()

	// seed data to the fake db
	courses = append(courses, Course{
		CourseId:   "123",
		CourseName: "Course1",
		Price:      11,
		Author: &Author{
			Fullname: "Auron Vila",
			Website:  "auronvila.com",
		},
	})

	courses = append(courses, Course{
		CourseId:   "321",
		CourseName: "Course2",
		Price:      22,
		Author: &Author{
			Fullname: "Auron Vila",
			Website:  "auronvila.com",
		},
	})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourseRecord).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", removeCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
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

	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		jsonResponse := fmt.Sprintf(`{"message": "Please send data", "courseName": "%s", "price": "%s"}`, newCourse.CourseName, newCourse.Price)
		err := json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			panic(err)
		}
		return
	}

	for _, course := range courses {
		if course.CourseName == newCourse.CourseName {
			json.NewEncoder(w).Encode("A course with this name exists")
			return
		}
	}

	newCourse.CourseId = strconv.Itoa(rand.Intn(1000))
	courses := append(courses, newCourse)
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
			json.NewEncoder(w).Encode("Could not find a course with given Id")
			return
		}
	}
}
