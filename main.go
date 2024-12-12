package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hectorhernandezalfonso/exercise_ms.git/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (optional, for development)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Retrieve MongoDB URI from environment variables
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	// Initialize MongoDB connection
	err = service.InitDatabase(uri)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Start the application
	app := service.New()
	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
}
