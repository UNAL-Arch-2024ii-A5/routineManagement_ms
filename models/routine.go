package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Routine struct {
	ID                primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	RoutineName       string               `json:"routine_name" bson:"routine_name"`
	RoutineDifficulty int                  `json:"routine_difficulty" bson:"routine_difficulty"`
	RoutineExercises  []primitive.ObjectID `json:"routine_exercises" bson:"routine_exercises"`
	RoutineMuscles    []MuscularGroup      `json:"routine_muscles" bson:"routine_muscles"`
	ImageUrl          string               `json:"image_url" bson:"image_url"`
	Owner             string               `json:"owner" bson:"owner"`
	Exercises         []Exercise           `json:"exercises" bson:"exercises"`
	CustomerId        []primitive.ObjectID `json:"customer_id" bson:"customer_id"` // New field
}
