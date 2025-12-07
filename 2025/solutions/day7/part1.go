// Package day7 solves puzzle available on https://adventofcode.com/2025/day/7
package day7

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Matrix(data)

	return numSplits(startX(in), 0, in)
}

func numSplits(x, yStart int, matrix [][]string) int {
	for y, row := range matrix[yStart:] {
		switch row[x] {
		case "|":
			return 0
		case ".":
			row[x] = "|"
		case "^":
			return 1 + numSplits(x-1, yStart+y, matrix) + numSplits(x+1, yStart+y, matrix)
		}
	}

	return 0
}
