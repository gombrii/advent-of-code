package day1

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
)

func Part1(file string) {
	in := input.Array(file)

	numStrings := findAll(in)
	sum := parseAndSum(numStrings)

	fmt.Println(sum)
}
