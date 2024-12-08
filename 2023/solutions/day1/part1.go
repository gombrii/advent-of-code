package day1

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2023", "day1", "part1", Part1)
}

func Part1(file string) {
	in := input.Slice(file)

	numStrings := findAll(in)
	sum := parseAndSum(numStrings)

	fmt.Println(sum)
}
