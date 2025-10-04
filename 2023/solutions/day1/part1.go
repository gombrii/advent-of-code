package day1

import (
	"github.com/gombrii/Advent-of-code/shared/input"
	"github.com/gombrii/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2023", "day1", "part1", Part1)
}

func Part1(file string) any {
	in := input.Slice(file)

	numStrings := findAll(in)
	sum := parseAndSum(numStrings)

	return sum
}
