package utils

import (
	"strings"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	if uuid == "" {
		t.Error("GenerateUUID returned an empty string")
	}
	if len(uuid) != 36 {
		t.Errorf("GenerateUUID returned a string of incorrect length: got %d, want 36", len(uuid))
	}
}

func TestGenerateUUIDRandom(t *testing.T) {
	uuid := GenerateUUIDRandom()
	if uuid == "" {
		t.Error("GenerateUUIDRandom returned an empty string")
	}
	if len(uuid) != 36 {
		t.Errorf("GenerateUUIDRandom returned a string of incorrect length: got %d, want 36", len(uuid))
	}
}

func TestGenerateCustomUUID(t *testing.T) {
	uid := "testuser123"
	uuid := GenerateCustomUUID(uid)
	parts := strings.Split(uuid, "-")

	if len(parts) != 3 {
		t.Errorf("GenerateCustomUUID returned incorrect format: got %s", uuid)
	}

	if !strings.Contains(uuid, uid) {
		t.Errorf("GenerateCustomUUID does not contain the provided uid: got %s, want to contain %s", uuid, uid)
	}

	if len(parts[2]) != 6 {
		t.Errorf("Random number part of CustomUUID is not 6 digits: got %s", parts[2])
	}
}
