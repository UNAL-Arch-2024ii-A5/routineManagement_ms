package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ExerciseName  string             `json:"exercise_name"`
	MuscularGroup []MuscularGroup    `json:"muscular_group"`
}

type MuscularGroup struct {
	MuscleID   uint64 `json:"muscle_id"`
	MuscleName string `json:"muscle_name"`
}
