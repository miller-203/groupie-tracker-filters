package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie/backend/handlers"
)

// Set the port
const port = ":8090"

// main is the entry point for the web server.
func main() {
	http.HandleFunc("/frontend/", handlers.Static)
	http.HandleFunc("/", handlers.Homepage)
	http.HandleFunc("/artist", handlers.ArtistPage)

	// Start the web server on port 8090. If the server fails to start, log the error and terminate.
	fmt.Println("Server listening at http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
