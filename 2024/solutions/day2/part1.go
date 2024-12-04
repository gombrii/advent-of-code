package day2

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day2", Part1)
}

func Part1(file string) {
	in := input.Slice(file)
	washed := wash(in)
	numSafe := numSafe(washed, safe)

	fmt.Println(numSafe)
}
