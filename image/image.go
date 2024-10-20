package utils

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/chai2010/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// SupportedFormats defines the formats supported for conversion.
var SupportedFormats = map[string]bool{
	"jpeg": true,
	"png":  true,
	"gif":  true,
	"bmp":  true,
	"tiff": true,
}

// ConvertToWebP converts an image (JPEG, PNG, GIF, BMP, TIFF) to WebP format.
func ConvertToWebP(input []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, errors.New("failed to decode input image: " + err.Error())
	}

	var buf bytes.Buffer
	options := &webp.Options{Lossless: false, Quality: 80}
	if err := webp.Encode(&buf, img, options); err != nil {
		return nil, errors.New("failed to encode image to WebP: " + err.Error())
	}

	return buf.Bytes(), nil
}

// ConvertFromWebP converts a WebP image to the specified format (JPEG, PNG, GIF, BMP, TIFF).
func ConvertFromWebP(input []byte, format string) ([]byte, error) {
	if !SupportedFormats[format] {
		return nil, errors.New("unsupported format: " + format)
	}

	img, err := webp.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, errors.New("failed to decode WebP image: " + err.Error())
	}

	var buf bytes.Buffer
	switch format {
	case "jpeg":
		options := &jpeg.Options{Quality: 85}
		if err := jpeg.Encode(&buf, img, options); err != nil {
			return nil, errors.New("failed to encode image to JPEG: " + err.Error())
		}
	case "png":
		if err := png.Encode(&buf, img); err != nil {
			return nil, errors.New("failed to encode image to PNG: " + err.Error())
		}
	case "gif":
		if err := gif.Encode(&buf, img, nil); err != nil {
			return nil, errors.New("failed to encode image to GIF: " + err.Error())
		}
	case "bmp":
		if err := bmp.Encode(&buf, img); err != nil {
			return nil, errors.New("failed to encode image to BMP: " + err.Error())
		}
	case "tiff":
		if err := tiff.Encode(&buf, img, nil); err != nil {
			return nil, errors.New("failed to encode image to TIFF: " + err.Error())
		}
	default:
		return nil, io.ErrUnexpectedEOF
	}

	return buf.Bytes(), nil
}
