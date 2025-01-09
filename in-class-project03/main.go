package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Student struct {
	ID      int
	Name    string
	Age     int
	Address Address
	Courses []Course
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Course struct {
	Code   string
	Name   string
	Credit int
}

type Studentlist struct {
	sync.Mutex
	Students map[int]Student
	NextID   int
}

func NewStore() *Studentlist {
	return &Studentlist{
		Students: make(map[int]Student),
		NextID:   1,
	}
}

func (store *Studentlist) Create(student Student) Student {
	store.Lock()
	defer store.Unlock()
	student.ID = store.NextID
	store.Students[student.ID] = student
	store.NextID++
	return student
}

func (store *Studentlist) GetAll() []Student {
	store.Lock()
	defer store.Unlock()
	all := make([]Student, 0, len(store.Students))
	for _, student := range store.Students {
		all = append(all, student)
	}
	return all
}

func (store *Studentlist) GetByID(id int) (Student, bool) {
	store.Lock()
	defer store.Unlock()
	student, found := store.Students[id]
	return student, found
}

func (store *Studentlist) Update(id int, student Student) (Student, bool) {
	store.Lock()
	defer store.Unlock()
	if _, found := store.Students[id]; found {
		student.ID = id
		store.Students[id] = student
		return student, true
	}
	return Student{}, false
}

func (store *Studentlist) Delete(id int) bool {
	store.Lock()
	defer store.Unlock()
	if _, found := store.Students[id]; found {
		delete(store.Students, id)
		return true
	}
	return false
}

var store = NewStore()

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var student Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		createdStudent := store.Create(student)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdStudent)
	case "GET":
		students := store.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/students/"):])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		var student Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		createdStudent := store.Create(student)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdStudent)
	case "GET":
		students := store.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	case "GETBYID":
		if student, found := store.GetByID(id); found {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	case "PUT":
		var student Student
		if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if updatedStudent, found := store.Update(id, student); found {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedStudent)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	case "DELETE":
		if store.Delete(id) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/students/", studentHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
