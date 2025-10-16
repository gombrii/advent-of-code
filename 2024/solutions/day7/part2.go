package day7

import (
	"fmt"
	"strconv"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

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

func Part2(data []byte) any {
	in := parse.Lines(data)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers, advancedManual)

	return total
}
