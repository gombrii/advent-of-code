package day2

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)
	washed := wash(in)
	numSafe := numSafe(washed, safeWithDampening)

	return numSafe
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
