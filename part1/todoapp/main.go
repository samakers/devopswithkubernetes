package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err) // Print the error
		return
	}

	// Get the port from the environment variable
	port := os.Getenv("PORT")

	// Check if the port is set
	if port == "" {
		fmt.Println("Error: PORT environment variable is not set.")
		return // Exit the application if PORT is not set
	}

	// Start the web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Output the server start message
	fmt.Printf("Server started on port %s\n", port)

	// Start the server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
