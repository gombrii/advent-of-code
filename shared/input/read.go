package input

import (
	"os"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
)

func Array(path string) []string {
	data := read(path)
	return strings.Split(string(data), "\n")
}

func String(path string) string {
	return string(read(path))
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
