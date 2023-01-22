package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/exercise", makeHTTPHandleFunc(s.handleExercise))

	router.HandleFunc("/workout", makeHTTPHandleFunc(s.handleWorkout))
	router.HandleFunc("/workout/{id}", makeHTTPHandleFunc(s.handleWorkout))

	router.HandleFunc("/userWorkouts", makeHTTPHandleFunc(s.handleGetUserWorkouts))

	fmt.Println("Server running on PORT: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleWorkout(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetWorkout(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateWorkout(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteWorkout(w, r)
	}

	return nil
}

func (s *APIServer) handleExercise(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAllExercises(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateExercise(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteExercise(w, r)
	}

	return nil
}

func (s *APIServer) handleGetWorkout(w http.ResponseWriter, r *http.Request) error {
	workoutId, _ := strconv.Atoi(mux.Vars(r)["id"])

	workout, err := s.storage.GetWorkout(workoutId)

	if err != nil {
		return nil
	}

	return WriteJSON(w, http.StatusOK, workout)
}

func (s *APIServer) handleGetUserWorkouts(w http.ResponseWriter, r *http.Request) error {
	workoutIdsDTO := new(WorkoutIdsDTO)
	if err := json.NewDecoder(r.Body).Decode(workoutIdsDTO); err != nil {
		return err
	}

	workouts, err := s.storage.GetUserWorkouts(workoutIdsDTO.WorkoutIds)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, workouts)
}

func (s *APIServer) handleCreateWorkout(w http.ResponseWriter, r *http.Request) error {
	workout := new(Workout)
	if err := json.NewDecoder(r.Body).Decode(workout); err != nil {
		return err
	}

	if err := s.storage.CreateWorkout(workout); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, workout)
}

func (s *APIServer) handleDeleteWorkout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateExercise(w http.ResponseWriter, r *http.Request) error {
	createExerciseDTO := new(CreateExerciseDTO)
	if err := json.NewDecoder(r.Body).Decode(createExerciseDTO); err != nil {
		return err
	}

	exercise := NewExercise(createExerciseDTO.Name, createExerciseDTO.Description, createExerciseDTO.Img)
	if err := s.storage.CreateExercise(exercise); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, exercise)
}

func (s *APIServer) handleDeleteExercise(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAllExercises(w http.ResponseWriter, r *http.Request) error {
	result, err := s.storage.GetAllExercises()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, result)
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
