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

var ErrDuplicateRoutineName = errors.New("routine with this name already exists")

type RoutineRepository struct {
	Collection *mongo.Collection
}

func NewRoutineRepository(db *mongo.Database) *RoutineRepository {
	return &RoutineRepository{
		Collection: db.Collection("routines"),
	}
}

func (r *RoutineRepository) checkDuplicateRoutineName(ctx context.Context, routineName string, excludeID *primitive.ObjectID) (bool, error) {
	normalizedName := strings.TrimSpace(strings.ToLower(routineName))
	pattern := primitive.Regex{Pattern: "^" + normalizedName + "$", Options: "i"}
	filter := bson.M{"routine_name": pattern}

	if excludeID != nil {
		filter["_id"] = bson.M{"$ne": excludeID}
	}

	count, err := r.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *RoutineRepository) CreateRoutine(ctx context.Context, routine models.Routine) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Check for duplicate routine name
	exists, err := r.checkDuplicateRoutineName(ctx, routine.RoutineName, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrDuplicateRoutineName
	}

	return r.Collection.InsertOne(ctx, routine)
}

func (r *RoutineRepository) GetRoutineByID(ctx context.Context, id primitive.ObjectID) (*models.Routine, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var routine models.Routine
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&routine)
	if err != nil {
		return nil, err
	}

	return &routine, nil
}

func (r *RoutineRepository) ListRoutines(ctx context.Context) ([]models.Routine, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var routines []models.Routine
	if err := cursor.All(ctx, &routines); err != nil {
		return nil, err
	}

	return routines, nil
}

func (r *RoutineRepository) UpdateRoutineByID(ctx context.Context, id primitive.ObjectID, routine models.Routine) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Check for duplicate routine name
	exists, err := r.checkDuplicateRoutineName(ctx, routine.RoutineName, &id)
	if err != nil {
		return err
	}
	if exists {
		return ErrDuplicateRoutineName
	}

	_, err = r.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": routine},
	)
	return err
}

func (r *RoutineRepository) DeleteRoutineByID(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
