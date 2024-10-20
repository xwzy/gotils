package utils

import (
	"bytes"
	"image"
	"image/png"
	"testing"

	"github.com/chai2010/webp"
)

func TestConvertToWebP(t *testing.T) {
	// Create a test image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		t.Fatalf("Failed to create test image: %v", err)
	}

	// Test conversion
	webpData, err := ConvertToWebP(buf.Bytes())
	if err != nil {
		t.Fatalf("ConvertToWebP failed: %v", err)
	}
	if len(webpData) == 0 {
		t.Error("ConvertToWebP returned empty data")
	}
}

func TestConvertFromWebP(t *testing.T) {
	// Create a test WebP image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	var webpBuf bytes.Buffer
	err := webp.Encode(&webpBuf, img, &webp.Options{Lossless: true})
	if err != nil {
		t.Fatalf("Failed to create test WebP image: %v", err)
	}

	for format := range SupportedFormats {
		t.Run(format, func(t *testing.T) {
			convertedData, err := ConvertFromWebP(webpBuf.Bytes(), format)
			if err != nil {
				t.Fatalf("ConvertFromWebP failed for %s: %v", format, err)
			}
			if len(convertedData) == 0 {
				t.Errorf("ConvertFromWebP returned empty data for %s", format)
			}
		})
	}
}

func TestConvertFromWebPUnsupportedFormat(t *testing.T) {
	_, err := ConvertFromWebP([]byte{}, "unsupported")
	if err == nil {
		t.Error("Expected error for unsupported format, got nil")
	}
}
