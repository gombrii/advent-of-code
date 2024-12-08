package day2

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day2", "part1", Part1)
}

func Part1(file string) {
	in := input.Slice(file)
	washed := wash(in)
	numSafe := numSafe(washed, safe)

	fmt.Println(numSafe)
}
