package day7

import (
	"fmt"
	"strconv"

	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part2("2024", "day7", Part2)
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

func Part2(file string) {
	in := input.Slice(file)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers, advancedManual)

	fmt.Println(total)
}
