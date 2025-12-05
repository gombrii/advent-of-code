// Package day5 solves puzzle available on https://adventofcode.com/2025/day/5
package day5

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)
	ranges, _ := compile(in)

	acc := 0
	for b, e := range ranges {
		acc += e - b + 1
	}

	return acc
}
