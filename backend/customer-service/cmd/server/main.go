package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Customer Service starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}

	log.Printf("Customer Service will listen on port %s", port)
	// Server initialization will be implemented in subsequent tasks
}
