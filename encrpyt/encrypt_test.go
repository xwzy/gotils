package utils

import (
	"testing"
)

func TestEncryptDecryptString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Short string", "Hello", "Hello"},
		{"Long string", "This is a longer string to test encryption and decryption", "This is a longer string to test encryption and decryption"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := EncryptString(tc.input)
			if err != nil {
				t.Fatalf("EncryptString failed: %v", err)
			}

			decrypted, err := DecryptString(encrypted)
			if err != nil {
				t.Fatalf("DecryptString failed: %v", err)
			}

			if decrypted != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, decrypted)
			}
		})
	}
}

func TestEncryptDecryptBytes(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{"Empty bytes", []byte{}, []byte{}},
		{"Short bytes", []byte("Hello"), []byte("Hello")},
		{"Long bytes", []byte("This is a longer byte slice to test encryption and decryption"), []byte("This is a longer byte slice to test encryption and decryption")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := EncryptBytes(tc.input)
			if err != nil {
				t.Fatalf("EncryptBytes failed: %v", err)
			}

			decrypted, err := DecryptBytes(encrypted)
			if err != nil {
				t.Fatalf("DecryptBytes failed: %v", err)
			}

			if string(decrypted) != string(tc.expected) {
				t.Errorf("Expected %q, but got %q", string(tc.expected), string(decrypted))
			}
		})
	}
}
