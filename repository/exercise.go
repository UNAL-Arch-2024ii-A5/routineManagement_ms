package repository

import (
	"context"
	"time"

	"github.com/hectorhernandezalfonso/exercise_ms.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExerciseRepository struct {
	Collection *mongo.Collection
}

// NewExerciseRepository initializes a new ExerciseRepository
func NewExerciseRepository(db *mongo.Database) *ExerciseRepository {
	return &ExerciseRepository{
		Collection: db.Collection("exercises"),
	}
}

func (r *ExerciseRepository) Create(ctx context.Context, exercise models.Exercise) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return r.Collection.InsertOne(ctx, exercise)
}

func (r *ExerciseRepository) List(ctx context.Context) ([]models.Exercise, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var exercises []models.Exercise
	if err := cursor.All(ctx, &exercises); err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *ExerciseRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Exercise, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var exercise models.Exercise
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&exercise)
	if err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (r *ExerciseRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, exercise models.Exercise) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": exercise})
	return err
}

func (r *ExerciseRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
