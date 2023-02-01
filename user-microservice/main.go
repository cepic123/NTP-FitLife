package main

import (
	"log"
)

func main() {
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3001", storage)
	server.Run()
}
