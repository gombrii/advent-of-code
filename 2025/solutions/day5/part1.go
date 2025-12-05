// Package day5 solves puzzle available on https://adventofcode.com/2025/day/5
package day5

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)
	ranges, ids := compile(in)

	acc := 0
	for _, id := range ids {
		if fresh(id, ranges) {
			acc++
		}
	}

	return acc
}

func fresh(id int, ranges map[int]int) bool {
	for b, e := range ranges {
		if id >= b && id <= e {
			return true
		}
	}

	return false
}
