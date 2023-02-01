package main

import (
	"log"
)

func main() {
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3003", storage)
	server.Run()
}
