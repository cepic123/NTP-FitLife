package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("YOLO")
	storage, err := NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", storage)
	server.Run()
}
