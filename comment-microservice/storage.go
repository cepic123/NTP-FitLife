package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StorageInterface interface {
	CreateComment(*Comment) error
	GetCommentBySubjectAndUser(int, int, string) (*Comment, error)
}

func (s *Storage) GetCommentBySubjectAndUser(userId, subjectId int, commentType string) (*Comment, error) {
	var comment Comment
	result := s.db.Where(&Comment{UserID: userId, SubjectID: subjectId, CommentType: commentType}).Find(&comment)

	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (s *Storage) CreateComment(comment *Comment) error {
	fmt.Println("CREATING COMMENT")
	if result := s.db.Create(comment); result.Error != nil {
		return result.Error
	}
	return nil
}

type Storage struct {
	db *gorm.DB
}

func NewStorage() (*Storage, error) {
	connStr := "user=postgres dbname=commentDB password=root sslmode=disable"
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
	db.AutoMigrate(&Comment{})

	return &Storage{
		db: db,
	}, nil
}
