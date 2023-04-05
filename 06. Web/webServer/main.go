package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>HOME PAGE</h1>\n"))
}

func HandlerStudents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>STUDENTS</h1>\n"))
}

func HandlerCourses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>FRONT + BACK</h1>\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/students", HandlerStudents)
	http.HandleFunc("/courses", HandlerCourses)

	log.Println("Start HTTP server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
