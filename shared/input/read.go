package input

import (
	"bytes"
	"os"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
)

func Slice(path string) []string {
	data := read(path)
	return strings.Split(string(data), "\n")
}

func String(path string) string {
	return string(read(path))
}

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
		exit.Errorf("invalid file path: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		exit.Errorf("could not read file: %v", err)
	}

	return data
}
