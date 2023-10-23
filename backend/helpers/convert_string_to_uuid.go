package helpers

import (
	"github.com/google/uuid"
)

// Function to parse UUID URL query parameters to UUID types
func ConvertStringToUUID(uuidString string) (uuid.UUID, error) {
	return uuid.Parse(uuidString)
}
