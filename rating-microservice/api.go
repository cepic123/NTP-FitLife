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

	router.HandleFunc("/rating", makeHTTPHandleFunc(s.handleRating))
	router.HandleFunc("/rating/{userId}/{workoutId}/{ratingType}", makeHTTPHandleFunc(s.handleGetRatingByUserAndSubject))

	fmt.Println("Server running on PORT: ", s.listenAddr)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)
	fmt.Println("HEREY")
	http.ListenAndServe(s.listenAddr, handler)
}

func (s *APIServer) handleGetRatingByUserAndSubject(w http.ResponseWriter, r *http.Request) error {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	workoutId, _ := strconv.Atoi(mux.Vars(r)["workoutId"])
	ratingType := mux.Vars(r)["ratingType"]

	rating, err := s.storage.GetRatingBySubjectAndUser(userId, workoutId, ratingType)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, rating)
}

func (s *APIServer) handleCreateRating(w http.ResponseWriter, r *http.Request) error {
	createRatingDTO := new(CreateRatingDTO)
	if err := json.NewDecoder(r.Body).Decode(createRatingDTO); err != nil {
		return err
	}

	rating, _ := s.storage.GetRatingBySubjectAndUser(createRatingDTO.UserID, createRatingDTO.SubjectID, createRatingDTO.RatingType)

	if rating.UserID != 0 {
		fmt.Println("USER ALREADY RATED")
		fmt.Println(rating)
		return WriteJSON(w, http.StatusOK, rating)
	}

	newRating := NewRating(createRatingDTO.UserID, createRatingDTO.SubjectID, createRatingDTO.Rating, createRatingDTO.Username, createRatingDTO.RatingType)
	if err := s.storage.CreateRating(newRating); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, newRating)
}

func (s *APIServer) handleUpdateRating(w http.ResponseWriter, r *http.Request) error {

	createRatingDTO := new(CreateRatingDTO)
	if err := json.NewDecoder(r.Body).Decode(createRatingDTO); err != nil {
		return err
	}

	rating, _ := s.storage.GetRatingBySubjectAndUser(createRatingDTO.UserID, createRatingDTO.SubjectID, createRatingDTO.RatingType)
	rating.Rating = createRatingDTO.Rating
	if rating.UserID == 0 {
		fmt.Println("rating DOESNT EXIST")
		fmt.Println(rating)
		return WriteJSON(w, http.StatusOK, rating)
	}

	if err := s.storage.UpdateRating(rating); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, rating)
}

func (s *APIServer) handleGetRating(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteRating(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleRating(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetRating(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateRating(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteRating(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateRating(w, r)
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
