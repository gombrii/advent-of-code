package day7

import (
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day7", "part1", Part1)
}

var simpleManual = manual{
	"*": func(a int, b int) int {
		return a * b
	},
	"+": func(a int, b int) int {
		return a + b
	},
}

func Part1(file string) any {
	in := input.Slice(file)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers, simpleManual)

	return total
}
