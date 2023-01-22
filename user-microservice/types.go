package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int           `json:"id" gorm:"primaryKey"`
	Email        string        `json:"email" gorm:"uniqueIndex"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	UserWorkouts []UserWorkout `json:"userWorkouts"`
}

type UserWorkout struct {
	gorm.Model
	ID                 int `json:"id" gorm:"primaryKey"`
	UserID             int `json:"userID"`
	WorkoutReferenceID int `json:"workoutReferenceID"`
}

func NewUser(username, password, email string) *User {
	return &User{
		Username:     username,
		Password:     password,
		Email:        email,
		UserWorkouts: []UserWorkout{},
	}
}

func NewUserWorkout(userId, workoutId int) *UserWorkout {
	return &UserWorkout{
		UserID:             userId,
		WorkoutReferenceID: workoutId,
	}
}
