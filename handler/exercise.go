package handler

import (
	"fmt"
	"net/http"
)

type Exercise struct{}

func (e *Exercise) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create exercise")
}

func (e *Exercise) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List exercises")
}

func (e *Exercise) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get exercise by ID")
}

func (e *Exercise) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update exercise by ID")
}

func (e *Exercise) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete exercise by ID")
}
