package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hectorhernandezalfonso/exercise_ms.git/handler"
	"github.com/hectorhernandezalfonso/exercise_ms.git/repository"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Exercise API!"))
	})

	// Initialize routes for exercises
	router.Route("/exercises", loadExerciseRoutes)
	// Initialize routes for routines
	router.Route("/routines", loadRoutineRoutes)

	return router
}

func loadExerciseRoutes(router chi.Router) {
	// Create a new instance of ExerciseRepository
	exerciseRepo := repository.NewExerciseRepository(DB.Client.Database("exercise_app"))
	// Pass the repository to the Exercise handler
	exerciseHandler := &handler.Exercise{Repo: exerciseRepo}

	// Map CRUD endpoints to the handler methods
	router.Post("/", exerciseHandler.CreateExercise)
	router.Get("/", exerciseHandler.ListExercises)
	router.Get("/{id}", exerciseHandler.GetExerciseByID)
	router.Put("/{id}", exerciseHandler.UpdateExerciseByID)
	router.Delete("/{id}", exerciseHandler.DeleteExerciseByID)
}

func loadRoutineRoutes(router chi.Router) {
	// Create instances of required repositories
	exerciseRepo := repository.NewExerciseRepository(DB.Client.Database("exercise_app"))
	routineRepo := repository.NewRoutineRepository(DB.Client.Database("exercise_app"))

	// Pass the repositories to the Routine handler
	routineHandler := &handler.Routine{
		Repo:         routineRepo,
		ExerciseRepo: exerciseRepo,
	}

	// Map CRUD endpoints to the handler methods
	router.Post("/", routineHandler.CreateRoutine)
	router.Get("/", routineHandler.ListRoutines)
	router.Get("/{id}", routineHandler.GetRoutineByID)
	router.Put("/{id}", routineHandler.UpdateRoutineByID)
	router.Delete("/{id}", routineHandler.DeleteRoutineByID)
}
