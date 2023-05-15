package filter

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Message struct {
	ID    string `json:"id"`
	Body  string `json:"body"`
	Image string `json:"image,omitempty"`
}

type BannedWordsResponse struct {
	Words []string `json:"words"`
}

type RejectedMessage struct {
	UID           string `json:"uid"`
	Contents      string `json:"message"`
	RejectionType string `json:"rejection_type"`
	Reason        string `json:"reason"`
}

func filterMessage(message *Message) (*Message, error) {
	// Check for a level 1 heading and at least one paragraph of normal text
	if !hasLevel1Heading(message.Body) || !hasParagraph(message.Body) {
		log.Printf("Rejected message with ID %s due to missing heading or paragraph", message.ID)
		rejectionType := "formatting"
		reason := "Rejected message due to missing heading or paragraph"
		err := writeRejectedMessage(message, rejectionType, reason)
		if err != nil {
			return nil, err
		}
		message.Body = reason
		return message, nil
	}
	// Check for inappropriate language
	if containsBannedWords(message.Body) {
		log.Printf("Rejected message with ID %s due to inappropriate language", message.ID)
		rejectionType := "language"
		reason := "Rejected message due to due to inappropriate language"
		err := writeRejectedMessage(message, rejectionType, reason)
		if err != nil {
			return nil, err
		}
		message.Body = reason
		return message, nil
	}
	// Handle images separately
	if message.Image != "" {
		// Store the message for image approval
		storeMessageForApproval(message)
		// Return the original message as is
		return message, nil
	}

	// Message passed all filters, return it
	return message, nil
}

func hasLevel1Heading(body string) bool {
	// Regular expression pattern for matching level 1 headings
	pattern := `^#\s.+`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Check if the body matches the pattern
	return regex.MatchString(body)
}

func hasParagraph(body string) bool {
	// Regular expression pattern for matching paragraphs
	pattern := `\n\n\s*.+`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(pattern)

	// Check if the body matches the pattern
	return regex.MatchString(body)
}

func containsBannedWords(body string) bool {
	// Fetch the banned words from the Language Service (you can replace this with your own implementation)
	bannedWords := fetchBannedWordsFromLanguageService()

	// Check if the message body contains any banned words
	// Return true if any banned words are found, false otherwise
	for _, word := range bannedWords {
		if strings.Contains(body, word) {
			return true
		}
	}

	return false
}

func fetchBannedWordsFromLanguageService() []string {
	// Return the fetched banned words
	response, err := http.Get("http://localhost:8080/banned-words")
	if err != nil {
		log.Printf("Failed to fetch banned words from Language Service: %v", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch banned words from Language Service. Status: %s", response.Status)
		return nil
	}

	var bannedWordsResponse BannedWordsResponse
	err = json.NewDecoder(response.Body).Decode(&bannedWordsResponse)
	if err != nil {
		log.Printf("Failed to decode banned words response: %v", err)
		return nil
	}

	return bannedWordsResponse.Words
}

func storeMessageForApproval(message *Message) {
	// Implement the logic to store the message for image approval
	// You can use a database, file storage, or any other suitable storage mechanism
	// Store the message along with its ID and image for later approval
}
