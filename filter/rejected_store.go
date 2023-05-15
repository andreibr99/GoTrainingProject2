package filter

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeRejectedMessage(message *Message, rejectionType, reason string) error {
	// Create a RejectedMessage struct with the necessary information
	rejectedMessage := RejectedMessage{
		UID:           message.ID,
		Contents:      message.Body,
		RejectionType: rejectionType,
		Reason:        reason,
	}

	// Read the existing rejected messages from the file
	fileData, err := os.ReadFile("rejection_store.json")
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read rejection store file: %v", err)
	}

	var rejectedMessages []RejectedMessage
	if len(fileData) > 0 {
		// Unmarshal the file data into a slice of RejectedMessage structs
		err = json.Unmarshal(fileData, &rejectedMessages)
		if err != nil {
			return fmt.Errorf("failed to unmarshal rejected messages: %v", err)
		}
	}

	// Append the new rejected message to the slice
	rejectedMessages = append(rejectedMessages, rejectedMessage)

	// Marshal the updated rejected messages into formatted JSON
	jsonData, err := json.MarshalIndent(rejectedMessages, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to marshal rejected messages: %v", err)
	}

	// Write the formatted JSON data to the file
	err = os.WriteFile("rejection_store.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write rejected messages to file: %v", err)
	}

	return nil
}
