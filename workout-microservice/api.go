package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router.HandleFunc("/exercise/{coachId}", makeHTTPHandleFunc(s.handleExercise))

	router.HandleFunc("/workout", makeHTTPHandleFunc(s.handleWorkout))
	router.HandleFunc("/workout/{id}", makeHTTPHandleFunc(s.handleWorkout))
	router.HandleFunc("/workout/rate/{id}/{rating}", makeHTTPHandleFunc(s.handleUpdateWorkoutRating))

	router.HandleFunc("/workout/calendar/{userId}/{workoutId}/{date}/{workoutName}", makeHTTPHandleFunc(s.handleAddCalendarEntry))

	router.HandleFunc("/workout/calendar/{id}", makeHTTPHandleFunc(s.handleCalendarEntries))
	router.HandleFunc("/workout/calendar/{id}", makeHTTPHandleFunc(s.handleCalendarEntries))

	router.HandleFunc("/userWorkouts", makeHTTPHandleFunc(s.handleGetUserWorkouts))

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

func (s *APIServer) handleCalendarEntries(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleCalendarEntriesForUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteCalendarEntry(w, r)
	}

	return nil
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

func (s *APIServer) handleDeleteCalendarEntry(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.storage.DeleteCalendarEntry(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleCalendarEntriesForUser(w http.ResponseWriter, r *http.Request) error {
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])

	calendarEntries, err := s.storage.GetCalendarEntriesForUser(userId)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, calendarEntries)
}

func (s *APIServer) handleAddCalendarEntry(w http.ResponseWriter, r *http.Request) error {
	workoutId, _ := strconv.Atoi(mux.Vars(r)["workoutId"])
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	date := mux.Vars(r)["date"]
	workoutName := mux.Vars(r)["workoutName"]

	calendarEntry := NewCalendarEntry(userId, workoutId, date, workoutName)

	if err := s.storage.CreateCalendarEntry(calendarEntry); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, calendarEntry)
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

func (s *APIServer) handleUpdateWorkoutRating(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	rating, _ := strconv.Atoi(mux.Vars(r)["rating"])

	workout, err := s.storage.GetWorkout(id)

	if err != nil {
		return err
	}

	workout.Rating = rating
	if workout.Name == "" {
		return WriteJSON(w, http.StatusOK, workout)
	}

	if err := s.storage.UpdateRating(workout); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, workout)
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
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.storage.DeleteWorkout(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleDeleteExercise(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.storage.DeleteExercise(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleCreateExercise(w http.ResponseWriter, r *http.Request) error {
	createExerciseDTO := new(CreateExerciseDTO)
	if err := json.NewDecoder(r.Body).Decode(createExerciseDTO); err != nil {
		return err
	}

	exercise := NewExercise(createExerciseDTO.Name, createExerciseDTO.Description, createExerciseDTO.Img, createExerciseDTO.CoachId)
	if err := s.storage.CreateExercise(exercise); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, exercise)
}

func (s *APIServer) handleGetAllExercises(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["coachId"])

	result, err := s.storage.GetAllExercises(id)

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
