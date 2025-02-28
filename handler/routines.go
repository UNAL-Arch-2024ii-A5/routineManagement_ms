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

type Routine struct {
	Repo         *repository.RoutineRepository
	ExerciseRepo *repository.ExerciseRepository
}

func (r *Routine) CreateRoutine(w http.ResponseWriter, req *http.Request) {
	var routine models.Routine
	if err := json.NewDecoder(req.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate difficulty range
	if routine.RoutineDifficulty < 1 || routine.RoutineDifficulty > 5 {
		http.Error(w, "Routine difficulty must be between 1 and 5", http.StatusBadRequest)
		return
	}

	// Validate and get exercises
	muscles := make([]models.MuscularGroup, 0)
	for _, exerciseID := range routine.RoutineExercises {
		exercise, err := r.ExerciseRepo.GetExerciseByID(req.Context(), exerciseID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid exercise ID: %s", exerciseID), http.StatusBadRequest)
			return
		}
		// Collect muscles from each exercise
		muscles = append(muscles, exercise.MuscularGroup...)
	}

	// Remove duplicate muscles
	routine.RoutineMuscles = removeDuplicateMuscles(muscles)

	ctx := req.Context()
	result, err := r.Repo.CreateRoutine(ctx, routine)
	if err != nil {
		if err == repository.ErrDuplicateRoutineName {
			http.Error(w, "A routine with this name already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (r *Routine) GetRoutineByID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid routine ID", http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	// Call the detailed version to get full exercise details
	routine, err := r.Repo.GetRoutineDetailedByID(ctx, objID)
	if err != nil {
		http.Error(w, "Routine not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routine)
}

func (r *Routine) ListRoutines(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	routines, err := r.Repo.ListRoutines(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch routines", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routines)
}

func (r *Routine) UpdateRoutineByID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid routine ID", http.StatusBadRequest)
		return
	}

	var routine models.Routine
	if err := json.NewDecoder(req.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate difficulty range
	if routine.RoutineDifficulty < 1 || routine.RoutineDifficulty > 5 {
		http.Error(w, "Routine difficulty must be between 1 and 5", http.StatusBadRequest)
		return
	}

	// Validate and get exercises
	muscles := make([]models.MuscularGroup, 0)
	for _, exerciseID := range routine.RoutineExercises {
		exercise, err := r.ExerciseRepo.GetExerciseByID(req.Context(), exerciseID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid exercise ID: %s", exerciseID), http.StatusBadRequest)
			return
		}
		muscles = append(muscles, exercise.MuscularGroup...)
	}

	// Remove duplicate muscles
	routine.RoutineMuscles = removeDuplicateMuscles(muscles)

	ctx := req.Context()
	if err := r.Repo.UpdateRoutineByID(ctx, objID, routine); err != nil {
		if err == repository.ErrDuplicateRoutineName {
			http.Error(w, "A routine with this name already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to update routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Routine updated successfully")
}

func (r *Routine) DeleteRoutineByID(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid routine ID", http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	if err := r.Repo.DeleteRoutineByID(ctx, objID); err != nil {
		http.Error(w, "Failed to delete routine", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Routine deleted successfully")
}

// Helper function to remove duplicate muscles
func removeDuplicateMuscles(muscles []models.MuscularGroup) []models.MuscularGroup {
	seen := make(map[uint64]bool)
	result := make([]models.MuscularGroup, 0)

	for _, muscle := range muscles {
		if !seen[muscle.MuscleID] {
			seen[muscle.MuscleID] = true
			result = append(result, muscle)
		}
	}
	return result
}
