package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
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

	router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLoginUser))
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/deleated", makeHTTPHandleFunc(s.handleGetDeleatedUsers))

	router.HandleFunc("/user/{id}", makeHTTPHandleFunc(s.handleUser))
	router.HandleFunc("/user/delete/{id}", makeHTTPHandleFunc(s.handleDeleteUserPerm))
	router.HandleFunc("/user/validate/{username}/{password}", makeHTTPHandleFunc(s.handleValidateUser))

	router.HandleFunc("/userWorkoutRefs/{id}", makeHTTPHandleFunc(s.handleGetUserWorkouts))

	router.HandleFunc("/user/restore/{id}", makeHTTPHandleFunc(s.handleRestoreUser))

	router.HandleFunc("/user/{userId}/{workoutId}", makeHTTPHandleFunc(s.handleWorkoutToUser))

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

func (s *APIServer) handleWorkoutToUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleAddWorkoutToUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleRemoveWorkoutFromUser(w, r)
	}

	return nil
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAllUsers(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}

	return nil
}

func (s *APIServer) handleGetUserWorkouts(w http.ResponseWriter, r *http.Request) error {
	userId, _ := strconv.Atoi(mux.Vars(r)["id"])

	userWorkouts, err := s.storage.GetUserWorkouts(userId)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, userWorkouts)
}

func (s *APIServer) handleValidateUser(w http.ResponseWriter, r *http.Request) error {
	username := mux.Vars(r)["username"]
	password := mux.Vars(r)["password"]

	user, err := s.storage.ValidateUser(username, password)

	if err != nil {
		return nil
	}

	if user == nil {
		return nil
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleRemoveWorkoutFromUser(w http.ResponseWriter, r *http.Request) error {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	workoutId, _ := strconv.Atoi(mux.Vars(r)["workoutId"])

	var result, err = s.storage.GetUserWorkout(userId, workoutId)

	empUser := UserWorkout{}
	if *result == empUser {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{"Pair doesn't exist in database"})
	}

	err = s.storage.DeleteUserWorkout(result.ID)

	return WriteJSON(w, http.StatusOK, err)
}

func (s *APIServer) handleAddWorkoutToUser(w http.ResponseWriter, r *http.Request) error {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	workoutId, _ := strconv.Atoi(mux.Vars(r)["workoutId"])

	var result, err = s.storage.GetUserWorkout(userId, workoutId)
	empUser := UserWorkout{}
	if *result != empUser {
		return WriteJSON(w, http.StatusInternalServerError, ApiError{"Pair already exists in database"})
	}

	userWorkout := NewUserWorkout(userId, workoutId)
	err = s.storage.CreateUserWorkout(userWorkout)

	return WriteJSON(w, http.StatusOK, err)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserDTO := new(CreateUserDTO)
	if err := json.NewDecoder(r.Body).Decode(createUserDTO); err != nil {
		return err
	}

	if createUserDTO.Role != "user" && createUserDTO.Role != "coach" {
		return WriteJSON(w, http.StatusInternalServerError, nil)
	}

	user := NewUser(createUserDTO.Username, createUserDTO.Password, createUserDTO.Email, createUserDTO.Role)
	if err := s.storage.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.storage.DeleteUser(id)

	return err
}

func (s *APIServer) handleDeleteUserPerm(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.storage.PermanentlyDeleteUser(id)

	return err
}

func (s *APIServer) handleRestoreUser(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.storage.RestoreUser(id)

	return err
}

func (s *APIServer) handleGetDeleatedUsers(w http.ResponseWriter, r *http.Request) error {
	result, err := s.storage.GetAllDeleatedUsers()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, result)
}

func (s *APIServer) handleGetAllUsers(w http.ResponseWriter, r *http.Request) error {
	result, err := s.storage.GetAllUsers()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, result)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	return WriteJSON(w, http.StatusOK, vars)
}

func (s *APIServer) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	loginUserDTO := new(LoginUserDTO)
	if err := json.NewDecoder(r.Body).Decode(loginUserDTO); err != nil {
		return err
	}

	user, err := s.storage.ValidateUser(loginUserDTO.Username, loginUserDTO.Password)

	if err != nil {
		return err
	}

	tokenString, err := createJWT(user)
	if err != nil {
		return err
	}

	loginResponseDTO := &LoginResponseDTO{
		ID:       user.ID,
		Username: user.Username,
		Token:    tokenString,
		Role:     user.Role,
	}
	return WriteJSON(w, http.StatusOK, loginResponseDTO)
}

func createJWT(user *User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"username":  user.Username,
		"password":  user.Password,
		"id":        user.ID,
		"role":      user.Role,
	}

	secret := "cepic"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
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
