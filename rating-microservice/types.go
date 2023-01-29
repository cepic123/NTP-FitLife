package main

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	UserID     int    `json:"userID"`
	Username   string `json:"username"`
	SubjectID  int    `json:"subjectID"`
	Rating     int    `json:"rating"`
	RatingType string `json:"ratingType"`
}

func NewRating(userID, subjectID, rating int, username, ratingType string) *Rating {
	return &Rating{
		UserID:     userID,
		Username:   username,
		SubjectID:  subjectID,
		Rating:     rating,
		RatingType: ratingType,
	}
}
