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
	DeleteCalendarEntry(int) error
	DeleteWorkout(int) error
	DeleteExercise(int) error
	GetCalendarEntriesForUser(int) (*[]CalendarEntry, error)
	CreateCalendarEntry(*CalendarEntry) error
}

func (s *Storage) DeleteWorkout(id int) error {
	result := s.db.Delete(&Workout{}, id)
	if result.Error != nil {
		return result.Error
	}

	result = s.db.Where("workout_id = ?", id).Delete(&CalendarEntry{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Storage) DeleteExercise(id int) error {
	result := s.db.Delete(&Exercise{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) DeleteCalendarEntry(id int) error {
	result := s.db.Delete(&CalendarEntry{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) GetCalendarEntriesForUser(userId int) (*[]CalendarEntry, error) {
	var calendarEntries []CalendarEntry

	result := s.db.Where(&CalendarEntry{UserId: userId}).Find(&calendarEntries)

	if result.Error != nil {
		return nil, result.Error
	}

	return &calendarEntries, nil
}

func (s *Storage) CreateCalendarEntry(calendarEntry *CalendarEntry) error {
	fmt.Println("Creating Calendar Entry")
	if result := s.db.Create(calendarEntry); result.Error != nil {
		return result.Error
	}
	return nil
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
	db.AutoMigrate(&CalendarEntry{})

	return &Storage{
		db: db,
	}, nil
}
