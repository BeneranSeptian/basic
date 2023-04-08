package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Model for courses - File
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper -File
func (course *Course) IsEmpty() bool {
	return course.CourseName == ""
}

func populateCourses() {
	courses = append(courses, Course{
		CourseId:    "001",
		CourseName:  "Belajar Golang",
		CoursePrice: 288,
		Author: &Author{
			Fullname: "Muhammad Septian",
			Website:  "septianbeneran.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "002",
		CourseName:  "Belajar Android Development",
		CoursePrice: 900,
		Author: &Author{
			Fullname: "Muhammad Septian",
			Website:  "septianbeneran.com",
		},
	})
}

func getCoursesGin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func createOneCourseGin(c *gin.Context) {
	//var newCourse Course
	//
	//error := c.ShouldBindJSON(&newCourse)
	//if error != nil {
	//	panic(error)
	//	return
	//}

	course := Course{
		CourseId:    "123678",
		CoursePrice: 9000,
		CourseName:  "Belajar Swift",
		Author:      &Author{Fullname: "Muhammad Septian", Website: "BeneranSeptian"}}

	courses = append(courses, course)
	c.IndentedJSON(http.StatusCreated, course)
}

func main() {
	//router := mux.NewRouter()
	router := gin.Default()
	populateCourses()

	// router
	//router.HandleFunc("/", serveHome).Methods(http.MethodGet)
	//router.HandleFunc("/courses", getAllCourses).Methods(http.MethodGet)
	//router.HandleFunc("/courses/{courseId}", getOneCourse).Methods(http.MethodGet)
	//router.HandleFunc("/courses/{courseId}", updateOneCourse).Methods(http.MethodPut)
	//router.HandleFunc("/courses", createOneCourse).Methods(http.MethodPost)
	//router.HandleFunc("/courses/{courseId}", deleteOneCourse).Methods(http.MethodDelete)

	router.GET("/courses", getCoursesGin)
	router.POST("/courses", createOneCourseGin)

	//listen and serve
	//log.Fatal(http.ListenAndServe(":4000", router))
	router.Run("localhost:8080")
}

// controllers - File

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get one id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["courseId"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	bodyReader := io.TeeReader(r.Body, &bytes.Buffer{})

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what if: body {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	err := json.NewDecoder(bodyReader).Decode(&course)
	if err != nil {
		log.Fatal(err)
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	// append new course to courses
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my id
	for index, course := range courses {
		if course.CourseId == params["courseId"] {
			courses = append(courses[:index], courses[:index+1]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["courseId"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO send a response when id not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my id
	for index, course := range courses {
		if course.CourseId == params["courseId"] {
			courses = append(courses[:index], courses[:index+1]...)
			json.NewEncoder(w).Encode("Deleted course with courseId: " + params["courseId"])
			return
		}
	}
}
