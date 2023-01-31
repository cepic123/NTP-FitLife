package main

type CreateExerciseDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
	CoachId     int    `json:"coachId"`
}

type WorkoutIdsDTO struct {
	WorkoutIds []int `json:"workoutIds"`
}
