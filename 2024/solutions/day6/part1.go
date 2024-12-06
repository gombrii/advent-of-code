package day6

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day6", Part1)
}

type guard struct {
	x     int
	y     int
	trace map[string]bool
}

func Part1(file string) {
	in := input.Matrix(file)

	fmt.Println("NOT IMPLEMENTED!")
}
