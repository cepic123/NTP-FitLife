package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type ApiError struct {
	Error string
}

func main() {
	fmt.Println("API GATEWAY")

	router := mux.NewRouter()

	// router.Use(authMiddleware)

	//USER MICROSERVICE
	router.HandleFunc("/login", redirect("http://localhost:3001"))
	router.HandleFunc("/user", redirect("http://localhost:3001"))
	router.HandleFunc("/user/{id}", redirect("http://localhost:3001"))

	//WORKOUT MICROSERVICE
	router.HandleFunc("/exercise", redirect("http://localhost:3002"))

	http.ListenAndServe(":3000", router)
}

func authMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling auth middleware")

		tokenStr := r.Header.Get("x-jwt-token")
		fmt.Println(tokenStr)
		_, err := validateJWT(tokenStr)

		fmt.Println("here")
		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}
		h.ServeHTTP(w, r)
	})
}

func validateJWT(tokenStr string) (*jwt.Token, error) {
	//TODO: PUT IN ENV
	secret := "cepic"
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func redirect(redirectAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		client := &http.Client{
			Timeout: time.Second * 10,
		}

		var req *http.Request
		var data []byte

		if r.Method == http.MethodPost {
			data, _ := ioutil.ReadAll(r.Body)
			req, _ = http.NewRequest(http.MethodPost, redirectAddr+r.URL.String(), bytes.NewBuffer(data))
		} else if r.Method == http.MethodGet {
			req, _ = http.NewRequest(http.MethodGet, redirectAddr+r.URL.String(), nil)
		} else if r.Method == http.MethodDelete {
			req, _ = http.NewRequest(http.MethodDelete, redirectAddr+r.URL.String(), nil)
		}

		res, err := client.Do(req)
		if err != nil {
			// log.Fatalln(err)
			return
		}
		defer res.Body.Close()

		data, _ = ioutil.ReadAll(res.Body)
		w.WriteHeader(res.StatusCode)
		w.Write([]byte(string(data)))
	}
}
