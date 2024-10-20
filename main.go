package main

import (
	"encoding/json"

	"github.com/xwzy/gotils/log"
)

func main() {
	TestLog()
}

func TestLog() {
	complexData := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"address": map[string]interface{}{
			"street": "123 Main St",
			"city":   "Anytown",
			"state":  "CA",
			"zip":    "12345",
		},
		"phoneNumbers": []string{
			"555-1234",
			"555-5678",
		},
		"isActive": true,
		"balance":  1234.56,
		"tags":     []string{"customer", "premium"},
		"metadata": map[string]interface{}{
			"lastLogin": "2023-04-01T14:30:00Z",
			"preferences": map[string]interface{}{
				"theme":      "dark",
				"newsletter": true,
			},
		},
	}

	jsonData, err := json.MarshalIndent(complexData, "", "  ")
	if err != nil {
		log.Error("Failed to marshal JSON:", err)
		return
	}

	log.Info("Complex JSON data:\n", string(jsonData))
}
