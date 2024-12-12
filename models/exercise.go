package models

type Exercise struct {
	ExerciseID    uint64          `json:"exercise_id"`
	ExerciseName  string          `json:"exercise_name"`
	MuscularGroup []MuscularGroup `json:"muscular_group"`
}

type MuscularGroup struct {
	MuscleID   uint64 `json:"muscle_id"`
	MuscleName string `json:"muscle_name"`
}
