package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ExerciseName  string             `json:"exercise_name" bson:"exercise_name"`
	ExerciseImage string             `json:"exercise_image" bson:"exercise_image"`
	ExerciseTime  int                `json:"exercise_time" bson:"exercise_time"` // in minutes
	ExerciseSets  int                `json:"exercise_sets" bson:"exercise_sets"`
	ExerciseReps  int                `json:"exercise_reps" bson:"exercise_reps"`
	MuscularGroup []MuscularGroup    `json:"muscular_group" bson:"muscular_group"`
}

type MuscularGroup struct {
	MuscleID   uint64 `json:"muscle_id" bson:"muscle_id"`
	MuscleName string `json:"muscle_name" bson:"muscle_name"`
}
