package main

import (
	_ "github.com/gomsim/Advent-of-code/2023/solutions/day1"
	_ "github.com/gomsim/Advent-of-code/2023/solutions/day2"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day1"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day10"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day2"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day3"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day4"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day5"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day6"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day7"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day8"
	_ "github.com/gomsim/Advent-of-code/2024/solutions/day9"

	"fmt"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

type input struct {
	Year string `arg:"-y,--year" default:"2024"`
	Day  string `arg:"-d,--day,required"`
	Part string `arg:"-p,--part,required"`
	In   string `arg:"-i,--input" default:"input"`
}

func main() {
	in := input{}
	arg.MustParse(&in)
	in.Day = fmt.Sprintf("day%s", in.Day)
	in.Part = fmt.Sprintf("part%s", in.Part)

	inputPath := fmt.Sprintf("%s/input/%s", in.Year, in.Day)
	inputFile := fmt.Sprintf("%s.txt", in.In)

	fmt.Printf("Running %s/%s/%s with %s\n", in.Year, in.Day, in.Part, inputFile)

	solution := registry.Registry[in.Year][in.Day][in.Part]
	if solution == nil {
		fmt.Println("Error: Solution not found!")
		os.Exit(1)
	}
	start := time.Now()
	result := solution(fmt.Sprintf("%s/%s", inputPath, inputFile))
	fmt.Println("Res:", result)
	fmt.Println("Dur:", time.Since(start))
}
