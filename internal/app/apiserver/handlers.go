package apiserver

import (
	"fmt"
	"github.com/zn11ch/SimpleWebApp/internal/model"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func (s *ApiServer) index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(filepath.Join("templates", "students", "index.html"))
	arr, _ := s.store.Student().ListAll()
	fmt.Println(arr)
	err := t.Execute(w, arr)
	if err != nil {
		log.Fatal("edit", err)
	}
}

func (s *ApiServer) add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles(filepath.Join("templates", "students", "add.html"))
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		fullName := strings.Join(r.Form["fullname"], "")
		faculty := strings.Join(r.Form["faculty"], "")
		course, _ := strconv.Atoi(strings.Join(r.Form["course"], ""))
		student := model.Student{FullName: fullName, Faculty: faculty, Course: course}

		_, err := s.store.Student().Create(&student)
		if err != nil {
			log.Fatal("edit", err)
		}
		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusFound)
	}
}

func (s *ApiServer) edit(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		studentId, _ := strconv.Atoi(r.URL.Path[len("/edit/"):])
		student, _ := s.store.Student().FindById(studentId)
		t, _ := template.ParseFiles(filepath.Join("templates", "students", "edit.html"))
		err := t.Execute(w, student)
		if err != nil {
			log.Fatal("edit", err)
		}
	case "POST":
		r.ParseForm()
		studentId, _ := strconv.Atoi(r.URL.Path[len("/edit/"):])
		student, _ := s.store.Student().FindById(studentId)
		fullName := strings.Join(r.Form["fullname"], "")
		faculty := strings.Join(r.Form["faculty"], "")
		course, _ := strconv.Atoi(strings.Join(r.Form["course"], ""))
		student.FullName = fullName
		student.Faculty = faculty
		student.Course = course
		_, err := s.store.Student().Update(student)
		if err != nil {
			log.Fatal("edit", err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (s *ApiServer) view(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		studentId, _ := strconv.Atoi(r.URL.Path[len("/view/"):])
		student, _ := s.store.Student().FindById(studentId)
		t, _ := template.ParseFiles(filepath.Join("templates", "students", "view.html"))
		err := t.Execute(w, student)
		if err != nil {
			log.Fatal("edit", err)
		}
	}
}
