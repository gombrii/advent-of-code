package day7

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

var simpleManual = manual{
	"*": func(a int, b int) int {
		return a * b
	},
	"+": func(a int, b int) int {
		return a + b
	},
}

func Part1(data []byte) any {
	in := parse.Lines(data)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers, simpleManual)

	return total
}
