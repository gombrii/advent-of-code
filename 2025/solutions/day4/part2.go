// Package day4 solves puzzle available on https://adventofcode.com/2025/day/4
package day4

import "github.com/gombrii/Advent-of-code/shared/parse"

func Part2(data []byte) any {
	in := parse.Matrix(data)

	acc := 0
	removed := true
	for removed {
		removed = false
		for y, row := range in {
			for x := range row {
				if accessible(in, x, y) {
					removed = true
					in[y][x] = "."
					acc++
				}
			}
		}
	}

	return acc
}
