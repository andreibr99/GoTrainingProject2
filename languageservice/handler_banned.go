package languageservice

import (
	"encoding/json"
	"net/http"
)

type BannedWordsResponse struct {
	Words []string `json:"words"`
}

func HandleBannedWords(w http.ResponseWriter, r *http.Request) {
	// Fetch the banned words/phrases from the store (you can replace this with your own data source)
	bannedWords := fetchBannedWords()

	// Create the response payload
	response := BannedWordsResponse{Words: bannedWords}

	// Encode the response payload to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.Write(jsonResponse)
}

func fetchBannedWords() []string {
	// This function can be implemented to fetch the banned words/phrases from a data source
	// For simplicity, we'll return a hardcoded list here
	return []string{
		"Harmony", "revolutionary", "bounce", "clue", "auction", "crew", "question",
		"flower", "rescue", "affair", "think", "night", "morale", "route", "regular",
		"veil", "ensure", "communication", "undertake", "gear", "professional",
		"judgment", "adult", "jaw", "death",
	}
}
