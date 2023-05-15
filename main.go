package main

import (
	"GoTrainingProject2/filter"
	"GoTrainingProject2/languageservice"
	"log"
	"net/http"
)

func main() {
	// Set up the HTTP routes and handlers
	http.HandleFunc("/banned-words", languageservice.HandleBannedWords)
	http.HandleFunc("/filter", filter.HandleFilterMessage)

	// Start the HTTP server
	log.Printf("Language Service listening on %s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
