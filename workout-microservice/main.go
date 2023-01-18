package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("WORKOUT MICROSERVICE")
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3002", storage)
	server.Run()
}
