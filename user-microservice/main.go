package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("NOLO")
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3001", storage)
	server.Run()
}
