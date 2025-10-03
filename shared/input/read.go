package input

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Slice returns the data in the specified file as a slice where each element represents a line of
// text.
func Slice(path string) []string {
	data := read(path)
	return strings.Split(string(data), "\n")
}

// String returns the data in the specified file as a continuous string.
func String(path string) string {
	return string(read(path))
}

// Deprecated: Use Matrix instead
func OldBadMatrix(path string) [][]byte {
	data := read(path)
	return bytes.Split(data, []byte("\n"))
}

func Matrix(path string) [][]string {
	data := read(path)
	matrix := make([][]string, 0)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix
}

func read(path string) []byte {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error: invalid file path: %v", err)
		os.Exit(1)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: could not read file: %v", err)
		os.Exit(1)
	}

	return data
}
