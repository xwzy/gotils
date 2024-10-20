package utils

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Compress(input []byte) ([]byte, error) {
	var buf bytes.Buffer

	// Create gzip writer with highest compression level
	gzipWriter, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	// Write input to gzip writer
	_, err = gzipWriter.Write(input)
	if err != nil {
		return nil, err
	}

	// Close the writer to flush any remaining data
	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decompress(input []byte) ([]byte, error) {
	// Create a bytes reader from the input
	reader := bytes.NewReader(input)

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Read decompressed data
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gzipReader)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
