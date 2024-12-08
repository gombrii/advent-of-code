package day2

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day2", "part2", Part2)
}

func Part2(file string) {
	in := input.Slice(file)
	washed := wash(in)
	numSafe := numSafe(washed, safeWithDampening)

	fmt.Println(numSafe)
}

func safeWithDampening(report []int) bool {
	if safe(report) {
		return true
	}
	for i := range report {
		if safe(append(report[:i:i], report[i+1:]...)) {
			return true
		}
	}
	return false
}
