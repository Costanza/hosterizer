package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Site Service starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8003"
	}

	log.Printf("Site Service will listen on port %s", port)
	// Server initialization will be implemented in subsequent tasks
}
