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
	return router
}

func loadExerciseRoutes(router chi.Router) {
	// Create a new instance of ExerciseRepository
	exerciseRepo := repository.NewExerciseRepository(DB.Client.Database("exercise_app"))

	// Pass the repository to the Exercise handler
	exerciseHandler := &handler.Exercise{Repo: exerciseRepo}

	// Map CRUD endpoints to the handler methods
	router.Post("/", exerciseHandler.Create)
	router.Get("/", exerciseHandler.List)
	router.Get("/{id}", exerciseHandler.GetByID)
	router.Put("/{id}", exerciseHandler.UpdateByID)
	router.Delete("/{id}", exerciseHandler.DeleteByID)
}
