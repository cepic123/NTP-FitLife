package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StorageInterface interface {
	CreateExercise(*Exercise) error
	GetAllExercises() (*[]Exercise, error)

	CreateWorkout(*Workout) error
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
	connStr := "user=postgres dbname=user-microservice password=root sslmode=disable"
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
	//TODO: PUT THIS SOMHERE ELSE
	db.AutoMigrate(&Workout{}, &Exercise{}, &Set{}, &Rep{})

	return &Storage{
		db: db,
	}, nil
}
