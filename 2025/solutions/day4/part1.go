// Package day4 solves puzzle available on https://adventofcode.com/2025/day/4
package day4

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Matrix(data, "")

	acc := 0
	for y, row := range in {
		for x := range row {
			if accessible(in, x, y) {
				acc++
			}
		}
	}

	return acc
}
