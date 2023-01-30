package main

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type Rep struct {
	gorm.Model
	ID         int `json:"id" gorm:"primaryKey"`
	SetID      int
	OrderNum   int `json:"orderNum"`
	NoReps     int `json:"noReps"`
	ExerciseID int
	Exercise   Exercise `json:"exercise"`
}

type Set struct {
	gorm.Model
	ID         int `json:"id" gorm:"primaryKey"`
	WorkoutID  int
	OrderNum   int   `json:"orderNum"`
	NoSets     int   `json:"noSets"`
	BreakLngth int   `json:"breakLngth"`
	Reps       []Rep `json:"reps"`
}

type Workout struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sets        []Set  `json:"sets"`
	Rating      int    `json:"rating"`
}

func NewExercise(name, description, img string) *Exercise {
	return &Exercise{
		Name:        name,
		Description: description,
		Img:         img,
	}
}
