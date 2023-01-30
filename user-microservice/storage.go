package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StorageInterface interface {
	CreateUser(*User) error
	CreateUserWorkout(*UserWorkout) error
	DeleteUser(int) error
	DeleteUserWorkout(int) error
	GetUser(int) (*User, error)
	GetUserWorkout(int, int) (*UserWorkout, error)
	GetUserWorkouts(int) (*[]UserWorkout, error)
	GetAllUsers() (*[]User, error)
	GetAllDeleatedUsers() (*[]User, error)
	ValidateUser(string, string) (*User, error)
	RestoreUser(int) error
	PermanentlyDeleteUser(int) error
}

type Storage struct {
	db *gorm.DB
}

func (s *Storage) CreateUserWorkout(userWorkout *UserWorkout) error {
	fmt.Println("CREATING REFERENCE")
	if result := s.db.Create(userWorkout); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) CreateUser(user *User) error {
	if result := s.db.Create(user); result.Error != nil {
		return result.Error
	}

	if user.Role == "coach" {
		if result := s.db.Unscoped().Model(&User{}).Where("id", user.ID).Update("deleted_at", user.CreatedAt); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s *Storage) DeleteUser(id int) error {
	result := s.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) PermanentlyDeleteUser(id int) error {
	result := s.db.Unscoped().Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) RestoreUser(id int) error {
	result := s.db.Unscoped().Model(&User{}).Where("id", id).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) DeleteUserWorkout(id int) error {
	result := s.db.Delete(&UserWorkout{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Storage) ValidateUser(username, password string) (*User, error) {
	var user User

	result := s.db.Where(&User{Username: username, Password: password}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *Storage) GetUser(id int) (*User, error) {
	var user = User{ID: id}

	result := s.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *Storage) GetAllUsers() (*[]User, error) {
	var users []User

	result := s.db.Unscoped().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (s *Storage) GetAllDeleatedUsers() (*[]User, error) {
	var users []User

	result := s.db.Unscoped().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (s *Storage) GetUserWorkouts(userId int) (*[]UserWorkout, error) {
	var userWorkouts []UserWorkout

	result := s.db.Where(&UserWorkout{UserID: userId}).Find(&userWorkouts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userWorkouts, nil
}

func (s *Storage) GetUserWorkout(userId, workoutId int) (*UserWorkout, error) {
	var userWorkouts UserWorkout

	result := s.db.Where(&UserWorkout{UserID: userId, WorkoutReferenceID: workoutId}).Find(&userWorkouts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userWorkouts, nil
}

func NewStorage() (*Storage, error) {
	connStr := "user=postgres dbname=userDB password=root sslmode=disable"
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
	db.AutoMigrate(&User{}, &UserWorkout{})

	return &Storage{
		db: db,
	}, nil
}
