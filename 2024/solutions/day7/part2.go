package day7

import (
	"fmt"
	"strconv"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/input"
	"github.com/gombrii/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day7", "part2", Part2)
}

var advancedManual = manual{
	"*": func(a int, b int) int {
		return a * b
	},
	"+": func(a int, b int) int {
		return a + b
	},
	"||": func(a int, b int) int {
		res, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		exit.If(err)
		return res
	},
}

func Part2(file string) any {
	in := input.Slice(file)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers, advancedManual)

	return total
}
