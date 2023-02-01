package main

import (
	"gorm.io/gorm"
)

type CalendarEntry struct {
	gorm.Model
	UserId      int    `json:"userId"`
	WorkoutId   int    `json:"workoutId"`
	Date        string `json:"date"`
	WorkoutName string `json:"workoutName"`
}

type Exercise struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
	CoachId     int    `json:"coachId"`
}

type Rep struct {
	gorm.Model
	SetID      int
	OrderNum   int `json:"orderNum"`
	NoReps     int `json:"noReps"`
	ExerciseID int
	Exercise   Exercise `json:"exercise"`
}

type Set struct {
	gorm.Model
	WorkoutID  int
	OrderNum   int   `json:"orderNum"`
	NoSets     int   `json:"noSets"`
	BreakLngth int   `json:"breakLngth"`
	Reps       []Rep `json:"reps"`
}

type Workout struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Sets        []Set  `json:"sets"`
	Rating      int    `json:"rating"`
	CoachId     int    `json:"coachId"`
}

func NewExercise(name, description, img string, coachId int) *Exercise {
	return &Exercise{
		Name:        name,
		Description: description,
		Img:         img,
		CoachId:     coachId,
	}
}

func NewCalendarEntry(userId, workoutId int, date, workoutName string) *CalendarEntry {
	return &CalendarEntry{
		UserId:      userId,
		WorkoutId:   workoutId,
		Date:        date,
		WorkoutName: workoutName,
	}
}
