package day7

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day7", "part1", Part1)
}

var simpleManual = manual{
	"*": func(a int, b int) int {
		return a * b
	},
	"+": func(a int, b int) int {
		return a + b
	},
}

func Part1(file string) {
	in := input.Slice(file)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers, simpleManual)

	fmt.Println(total)
}
