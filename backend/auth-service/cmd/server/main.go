package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Auth Service starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	log.Printf("Auth Service will listen on port %s", port)
	// Server initialization will be implemented in subsequent tasks
}
