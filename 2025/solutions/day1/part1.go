// Package day1 solves puzzle available on https://adventofcode.com/2025/day/1
package day1

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)
	parsed := parseData(in)

	curr := 50
	acc := 0
	for _, n := range parsed {
		curr = rotate(curr, n)
		if curr == 0 {
			acc++
		}
	}

	return acc
}

func rotate(curr, shift int) int {
	return ((curr+shift)%100 + 100) % 100
}
