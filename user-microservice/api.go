package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUser))

	router.HandleFunc("/user/{id}", makeHTTPHandleFunc(s.handleGetUser))

	fmt.Println("Server running on PORT: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
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

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserDTO := new(CreateUserDTO)
	if err := json.NewDecoder(r.Body).Decode(createUserDTO); err != nil {
		return err
	}

	user := NewUser(createUserDTO.Username, createUserDTO.Password)
	if err := s.storage.CreateUser(user); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
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

	// account := NewUser("ACA", "PASSWORD")

	return WriteJSON(w, http.StatusOK, vars)
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
