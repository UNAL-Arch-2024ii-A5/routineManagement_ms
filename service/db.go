package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	ExerciseDB *mongo.Collection
}

var DB *Database

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

func InitDatabase(uri string) error {
	client, err := ConnectMongoDB(uri)
	if err != nil {
		return err
	}

	// Initialize specific collections
	exerciseDB := client.Database("exercise_app").Collection("exercises")

	DB = &Database{
		Client:     client,
		ExerciseDB: exerciseDB,
	}
	return nil
}
