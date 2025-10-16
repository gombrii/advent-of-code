package day2

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)
	washed := wash(in)
	numSafe := numSafe(washed, safe)

	return numSafe
}
