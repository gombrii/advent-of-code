package day1

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)

	numStrings := findAll(in)
	sum := parseAndSum(numStrings)

	return sum
}
