package parse

import (
	"bytes"
	"strings"
)

// Lines converts data into text and returns it as a slice where each element represents a line.
func Lines(data []byte) []string {
	return strings.Split(string(data), "\n")
}

// String returns data as a continuous string.
func String(data []byte) string {
	return string(data)
}

// Deprecated: Use Matrix instead
func OldBadMatrix(data []byte) [][]byte {
	return bytes.Split(data, []byte("\n"))
}

func Matrix(data []byte, delimiter string) [][]string {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	matrix := make([][]string, len(lines))

	for i, line := range lines {
		matrix[i] = strings.Split(line, delimiter)
	}

	return matrix
}