package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Ecommerce Service starting...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8006"
	}

	log.Printf("Ecommerce Service will listen on port %s", port)
	// Server initialization will be implemented in subsequent tasks
}
