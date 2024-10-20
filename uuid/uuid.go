package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// GenerateUUID generates a UUID based on the current timestamp
func GenerateUUID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		// handle error appropriately
		return ""
	}
	return u.String()
}

// GenerateUUIDv4 generates a random UUID (version 4)
func GenerateUUIDRandom() string {
	u, err := uuid.NewRandom()
	if err != nil {
		// handle error appropriately
		return ""
	}
	return u.String()
}

// GenerateCustomUUID generates a custom UUID using timestamp, uid, and a random number
func GenerateCustomUUID(uid string) string {
	timestamp := time.Now().UnixMilli()
	randomNumber := rand.Intn(1000000)
	customUUID := fmt.Sprintf("%d-%s-%06d", timestamp, uid, randomNumber)
	return customUUID
}
