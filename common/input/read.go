package input

import (
	"os"
	"strings"

	"github.com/gomsim/Advent-of-code/common/exit"
)

func Array() []string {
	data := read(inputFile())
	return strings.Split(string(data), "\n")
}

func inputFile() string {
	if len(os.Args) < 2 {
		exit.Error("no input file provided!")
	}
	return os.Args[1]
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
