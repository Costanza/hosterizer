package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Policy Service starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8005"
	}

	log.Printf("Policy Service will listen on port %s", port)
	// Server initialization will be implemented in subsequent tasks
}
