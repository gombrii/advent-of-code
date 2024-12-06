package main

import (
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day6"
	"fmt"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day3"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day4"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day5"

	"github.com/alexflint/go-arg"
	_ "github.com/gomsim/Advent-of-code/2023/solutions/day1"
	_ "github.com/gomsim/Advent-of-code/2023/solutions/day2"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day1"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day2"
	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/register"
)

type input struct {
	Year string `arg:"-y,--year" default:"2024"`
	Day  string `arg:"required, positional"`
	Part string `arg:"required, positional"`
	In   string `arg:"required, positional"`
}

func main() {
	in := input{}
	arg.MustParse(&in)
	in.Day = fmt.Sprintf("day%s", in.Day)
	in.Part = fmt.Sprintf("part%s", in.Part)

	inputPath := fmt.Sprintf("%s/input/%s", in.Year, in.Day)
	inputFile := fmt.Sprintf("%s.txt", in.In)

	fmt.Printf("Running %s/%s/%s with %s\n", in.Year, in.Day, in.Part, inputFile)

	solution := register.Registry[in.Year][in.Day][in.Part]
	if solution == nil {
		exit.Errorf("Solution not found!")
	}
	solution(fmt.Sprintf("%s/%s", inputPath, inputFile))
}
