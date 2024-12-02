package main

import (
	"fmt"
	"os"

	_ "github.com/gomsim/Advent-of-code/2023/solutions/day1"
	_ "github.com/gomsim/Advent-of-code/2023/solutions/day2"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day1"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day2"
	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/register"
)

const (
	year = "2024"

	dayIndex   = 1
	partIndex  = 2
	inputIndex = 3
)

func main() {
	if len(os.Args) != 4 {
		exit.Errorf("Need 3 arguments (daynumber, partnumber, input filename), got %d", len(os.Args)-1)
	}

	day := fmt.Sprintf("day%s", os.Args[dayIndex])
	part := fmt.Sprintf("part%s", os.Args[partIndex])
	inputPath := fmt.Sprintf("%s/input/%s", year, day)
	inputFile := fmt.Sprintf("%s.txt", os.Args[inputIndex])

	fmt.Printf("Running %s/%s/%s with %s\n", year, day, part, inputFile)

	solution := register.Registry[year][day][part]
	if solution == nil {
		exit.Errorf("Solution not found!")
	}
	solution(fmt.Sprintf("%s/%s", inputPath, inputFile))
}
