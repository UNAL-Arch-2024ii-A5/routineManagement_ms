package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/hectorhernandezalfonso/exercise_ms.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrDuplicateExerciseName = errors.New("exercise with this name already exists")

type ExerciseRepository struct {
	Collection *mongo.Collection
}

func NewExerciseRepository(db *mongo.Database) *ExerciseRepository {
	return &ExerciseRepository{
		Collection: db.Collection("exercises"),
	}
}

func normalizeExerciseName(name string) string {
	// Convert to lowercase and trim spaces
	normalized := strings.TrimSpace(strings.ToLower(name))
	// Remove extra spaces between words
	return strings.Join(strings.Fields(normalized), " ")
}

// checkDuplicateExerciseName checks if an exercise with the given name already exists (case-insensitive)
func (r *ExerciseRepository) checkDuplicateExerciseName(ctx context.Context, exerciseName string, excludeID *primitive.ObjectID) (bool, error) {
	normalizedName := normalizeExerciseName(exerciseName)

	// Create a case-insensitive regex pattern
	pattern := primitive.Regex{Pattern: "^" + normalizedName + "$", Options: "i"}
	filter := bson.M{"exercisename": pattern}

	if excludeID != nil {
		filter["_id"] = bson.M{"$ne": excludeID}
	}

	count, err := r.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ExerciseRepository) CreateExercise(ctx context.Context, exercise models.Exercise) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Normalize the exercise name before saving
	exercise.ExerciseName = strings.TrimSpace(exercise.ExerciseName)

	// Check for duplicate exercise name
	exists, err := r.checkDuplicateExerciseName(ctx, exercise.ExerciseName, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicateExerciseName
	}

	return r.Collection.InsertOne(ctx, exercise)
}

func (r *ExerciseRepository) ListExercises(ctx context.Context) ([]models.Exercise, error) {
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

func (r *ExerciseRepository) GetExerciseByID(ctx context.Context, id primitive.ObjectID) (*models.Exercise, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var exercise models.Exercise
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&exercise)
	if err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (r *ExerciseRepository) UpdateExerciseByID(ctx context.Context, id primitive.ObjectID, exercise models.Exercise) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Normalize the exercise name before updating
	exercise.ExerciseName = strings.TrimSpace(exercise.ExerciseName)

	// Check for duplicate exercise name, excluding the current exercise
	exists, err := r.checkDuplicateExerciseName(ctx, exercise.ExerciseName, &id)
	if err != nil {
		return err
	}
	if exists {
		return ErrDuplicateExerciseName
	}

	_, err = r.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": exercise},
	)
	return err
}

func (r *ExerciseRepository) DeleteExerciseByID(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
