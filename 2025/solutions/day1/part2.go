// Package day1 solves puzzle available on https://adventofcode.com/2025/day/1
package day1

import (
	"math"

	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)
	parsed := parseData(in)

	curr := 50
	acc := 0
	for _, n := range parsed {
		new, passes := numPasses(curr, n)
		acc += passes
		curr = new
	}

	return acc
}

func numPasses(old, shift int) (new, passes int) {
	new = old + shift

	passes = int(math.Abs(float64(new / 100)))
	if old != 0 && new <= 0 {
		passes++
	}

	return (new%100 + 100) % 100, passes
}
