package filter

import (
	"fmt"
	"time"
)

func GenerateUniqueID() string {
	// We are using a simple timestamp-based ID
	return fmt.Sprintf("%d", getCurrentTimestamp())
}

func getCurrentTimestamp() int64 {
	// Get the current time in UTC
	now := time.Now().UTC()

	// Calculate the timestamp in milliseconds
	timestamp := now.UnixNano() / int64(time.Millisecond)
	// Sleep to ensure different IDs
	time.Sleep(1 * time.Microsecond)
	return timestamp
}
