package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("API GATEWAY")

	router := mux.NewRouter()

	router.HandleFunc("/user", redirect("http://localhost:3001"))
	router.HandleFunc("/user/{id}", redirect("http://localhost:3001"))

	http.ListenAndServe(":3000", router)
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
