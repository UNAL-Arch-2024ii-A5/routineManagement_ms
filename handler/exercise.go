package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hectorhernandezalfonso/exercise_ms.git/models"
	"github.com/hectorhernandezalfonso/exercise_ms.git/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	Repo *repository.ExerciseRepository
}

func (e *Exercise) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creando")
	var exercise models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	result, err := e.Repo.Create(ctx, exercise)
	if err != nil {
		http.Error(w, "Failed to create exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Created exercise with ID: %v", result)
}

func (e *Exercise) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Devolviendo")
	ctx := r.Context()
	exercises, err := e.Repo.List(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch exercises", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

func (e *Exercise) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get ID")
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	exercise, err := e.Repo.GetByID(ctx, objID)
	if err != nil {
		http.Error(w, "Exercise not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercise)
}

func (e *Exercise) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Actualizando")
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	var exercise models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := e.Repo.UpdateByID(ctx, objID, exercise); err != nil {
		http.Error(w, "Failed to update exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Exercise updated successfully")
}

func (e *Exercise) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Borrando")
	id := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := e.Repo.DeleteByID(ctx, objID); err != nil {
		http.Error(w, "Failed to delete exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Exercise deleted successfully")
}
