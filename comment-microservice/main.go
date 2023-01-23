package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("COMMENT MICROSERVICE")
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3003", storage)
	server.Run()
}
