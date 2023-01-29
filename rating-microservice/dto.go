package main

type CreateRatingDTO struct {
	UserID     int    `json:"userID"`
	Username   string `json:"username"`
	SubjectID  int    `json:"subjectID"`
	Rating     int    `json:"rating"`
	RatingType string `json:"ratingType"`
}
