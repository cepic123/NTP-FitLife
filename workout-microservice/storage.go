package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StorageInterface interface {
	CreateExercise(*Exercise) error
	GetAllExercises() (*[]Exercise, error)

	GetWorkout(int) (*Workout, error)
	GetUserWorkouts([]int) (*[]Workout, error)
	CreateWorkout(*Workout) error
	UpdateRating(*Workout) error
}

func (s *Storage) CreateExercise(exercise *Exercise) error {
	fmt.Println("IN STORAGE")
	if result := s.db.Create(exercise); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) CreateWorkout(workout *Workout) error {
	fmt.Println("IN STORAGE")
	if result := s.db.Create(workout); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) UpdateRating(workout *Workout) error {
	fmt.Println("UPDATING WORKOUT")
	if result := s.db.Save(workout); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) GetUserWorkouts(workoutIds []int) (*[]Workout, error) {
	var workouts []Workout

	result := s.db.Find(&workouts, workoutIds)

	if result.Error != nil {
		return nil, result.Error
	}
	return &workouts, nil
}

func (s *Storage) GetWorkout(workoutId int) (*Workout, error) {
	var workout Workout

	result := s.db.Preload("Sets.Reps.Exercise").Preload(clause.Associations).Find(&workout, workoutId)

	if result.Error != nil {
		return nil, result.Error
	}
	return &workout, nil
}

func (s *Storage) GetAllExercises() (*[]Exercise, error) {
	var exercises []Exercise

	result := s.db.Find(&exercises)
	if result.Error != nil {
		return nil, result.Error
	}
	return &exercises, nil
}

type Storage struct {
	db *gorm.DB
}

func NewStorage() (*Storage, error) {
	connStr := "user=postgres dbname=workoutDB password=root sslmode=disable"
	PgDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := PgDB.Ping(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: PgDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Workout{}, &Exercise{}, &Set{}, &Rep{})

	return &Storage{
		db: db,
	}, nil
}
