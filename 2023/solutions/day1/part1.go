package day1

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2023", "day1", Part1)
}

func Part1(file string) {
	in := input.Slice(file)

	numStrings := findAll(in)
	sum := parseAndSum(numStrings)

	fmt.Println(sum)
}
