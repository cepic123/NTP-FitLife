package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int           `json:"id" gorm:"primaryKey"`
	Email        string        `json:"email" gorm:"uniqueIndex"`
	Username     string        `json:"username" gorm:"uniqueIndex"`
	Password     string        `json:"password"`
	Role         string        `json:"role"`
	UserWorkouts []UserWorkout `json:"userWorkouts"`
}

type UserWorkout struct {
	gorm.Model
	ID                 int `json:"id" gorm:"primaryKey"`
	UserID             int `json:"userID"`
	WorkoutReferenceID int `json:"workoutReferenceID"`
}

func NewUser(username, password, email, role string) *User {
	return &User{
		Username:     username,
		Password:     password,
		Email:        email,
		Role:         role,
		UserWorkouts: []UserWorkout{},
	}
}

func NewUserWorkout(userId, workoutId int) *UserWorkout {
	return &UserWorkout{
		UserID:             userId,
		WorkoutReferenceID: workoutId,
	}
}
