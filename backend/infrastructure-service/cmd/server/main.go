package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Infrastructure Service starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8004"
	}

	log.Printf("Infrastructure Service will listen on port %s", port)
	// Server initialization will be implemented in subsequent tasks
}
