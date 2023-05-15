package filter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleFilterMessage(w http.ResponseWriter, r *http.Request) {
	// Example message data
	//messageData := `{
	//	"body": "# Simple Message\n\nThis is a good message."
	//}`

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid request method")
		return
	}

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	// Generate a unique ID for the message
	message.ID = GenerateUniqueID()

	// Perform the message filtering
	filteredMessage, err := filterMessage(&message)
	if err != nil {
		http.Error(w, "Failed to filter message", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Encode the filtered message to JSON
	jsonResponse, err := json.Marshal(filteredMessage)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.Write(jsonResponse)
}
