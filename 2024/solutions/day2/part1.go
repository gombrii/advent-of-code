package day2

import (
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day2", "part1", Part1)
}

func Part1(file string) any {
	in := input.Slice(file)
	washed := wash(in)
	numSafe := numSafe(washed, safe)

	return numSafe
}
