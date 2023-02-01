package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/comment", makeHTTPHandleFunc(s.handleComment))
	router.HandleFunc("/comment/{workoutId}/{commentType}", makeHTTPHandleFunc(s.handleGetCommentsBySubject))
	router.HandleFunc("/comment/{userId}/{workoutId}/{commentType}", makeHTTPHandleFunc(s.handleGetCommentByUserAndSubject))

	fmt.Println("Server running on PORT: ", s.listenAddr)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)
	http.ListenAndServe(s.listenAddr, handler)
}

func (s *APIServer) handleGetCommentsBySubject(w http.ResponseWriter, r *http.Request) error {
	workoutId, _ := strconv.Atoi(mux.Vars(r)["workoutId"])
	commentType := mux.Vars(r)["commentType"]

	comments, err := s.storage.GetCommentsBySubject(workoutId, commentType)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, comments)
}

func (s *APIServer) handleGetCommentByUserAndSubject(w http.ResponseWriter, r *http.Request) error {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	workoutId, _ := strconv.Atoi(mux.Vars(r)["workoutId"])
	commentType := mux.Vars(r)["commentType"]

	comment, err := s.storage.GetCommentBySubjectAndUser(userId, workoutId, commentType)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, comment)
}

func (s *APIServer) handleCreateComment(w http.ResponseWriter, r *http.Request) error {

	createCommentDTO := new(CreateCommentDTO)
	if err := json.NewDecoder(r.Body).Decode(createCommentDTO); err != nil {
		return err
	}

	comment, _ := s.storage.GetCommentBySubjectAndUser(createCommentDTO.UserID, createCommentDTO.SubjectID, createCommentDTO.CommentType)

	if comment.Comment != "" {
		return WriteJSON(w, http.StatusOK, comment)
	}

	user := NewComment(createCommentDTO.UserID, createCommentDTO.SubjectID, createCommentDTO.Username, createCommentDTO.Comment, createCommentDTO.CommentType)
	if err := s.storage.CreateComment(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleUpdateComment(w http.ResponseWriter, r *http.Request) error {

	createCommentDTO := new(CreateCommentDTO)
	if err := json.NewDecoder(r.Body).Decode(createCommentDTO); err != nil {
		return err
	}

	comment, _ := s.storage.GetCommentBySubjectAndUser(createCommentDTO.UserID, createCommentDTO.SubjectID, createCommentDTO.CommentType)
	comment.Comment = createCommentDTO.Comment
	if comment.Comment == "" {
		return WriteJSON(w, http.StatusOK, comment)
	}

	if err := s.storage.UpdateComment(comment); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, comment)
}

func (s *APIServer) handleDeleteComment(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetComment(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleComment(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetComment(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateComment(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteComment(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateComment(w, r)
	}
	return nil
}

type APIServer struct {
	listenAddr string
	storage    StorageInterface
}

func NewAPIServer(listenAddr string, storage StorageInterface) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
