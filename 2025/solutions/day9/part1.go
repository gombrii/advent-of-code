// Package day9 solves puzzle available on https://adventofcode.com/2025/day/9
package day9

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)
	tiles := parseData(in)

	return biggestRect(tiles)
}

func biggestRect(tiles [][2]int) int {
	hi := 0
	for _, tile := range tiles {
		for _, other := range tiles {
			area := area(tile, other)
			if area > hi {
				hi = area
			}
		}
	}

	return hi
}
