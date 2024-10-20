package utils

import (
	"bytes"
	"testing"
)

func TestCompressDecompressBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input []byte
	}{
		{"Empty input", []byte{}},
		{"Short input", []byte("Hello, World!")},
		{"Long input", bytes.Repeat([]byte("Lorem ipsum dolor sit amet. "), 100)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			compressed, err := Compress(tc.input)
			if err != nil {
				t.Fatalf("Compress failed: %v", err)
			}

			decompressed, err := Decompress(compressed)
			if err != nil {
				t.Fatalf("Decompress failed: %v", err)
			}

			if !bytes.Equal(decompressed, tc.input) {
				t.Errorf("Decompressed data does not match original input")
			}

			if len(compressed) > 0 && len(compressed) >= len(tc.input) {
				t.Logf("Compression ratio: %.2f", float64(len(compressed))/float64(len(tc.input)))
			}
		})
	}
}
