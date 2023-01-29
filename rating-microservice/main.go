package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("RATING MICROSERVICE")
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3004", storage)
	server.Run()
}
