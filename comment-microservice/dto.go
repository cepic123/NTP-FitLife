package main

type CreateCommentDTO struct {
	UserID      int    `json:"userID"`
	Username    string `json:"username"`
	SubjectID   int    `json:"subjectID"`
	Comment     string `json:"comment"`
	CommentType string `json:"commentType"`
}
