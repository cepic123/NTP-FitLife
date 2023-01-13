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
	DeleteUser(int) error
	GetUser(int) (*User, error)
	GetAllUsers() (*[]User, error)
}

type Storage struct {
	db *gorm.DB
}

func (s *Storage) CreateUser(user *User) error {
	fmt.Println("IN STORAGE")
	if result := s.db.Create(user); result.Error != nil {
		return result.Error
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

func (s *Storage) GetUser(int) (*User, error) {
	return nil, nil
}

func (s *Storage) GetAllUsers() (*[]User, error) {
	var users []User

	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
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
	db.AutoMigrate()

	return &Storage{
		db: db,
	}, nil
}
