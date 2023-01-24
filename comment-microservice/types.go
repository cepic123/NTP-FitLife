package main

import (
	"gorm.io/gorm"
)

// TODO: BRISI ID OD SVUDA
type Comment struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      int    `json:"userID"`
	Username    string `json:"username"`
	SubjectID   int    `json:"subjectID"`
	Comment     string `json:"comment"`
	CommentType string `json:"commentType"`
}

func NewComment(userID, subjectID int, username, comment, commentType string) *Comment {
	return &Comment{
		UserID:      userID,
		Username:    username,
		SubjectID:   subjectID,
		Comment:     comment,
		CommentType: commentType,
	}
}
