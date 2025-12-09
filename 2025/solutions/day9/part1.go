// Package day9 solves puzzle available on https://adventofcode.com/2025/day/9
package day9

import (
	"math"
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)
	tiles := parseData(in)

	return biggestRect(tiles)
}

func parseData(data []string) [][]int {
	tiles := make([][]int, len(data))

	for i, line := range data {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		exit.If(err)
		y, err := strconv.Atoi(parts[1])
		exit.If(err)

		tiles[i] = []int{x, y}
	}

	return tiles
}

func biggestRect(tiles [][]int) int {
	hi := 0
	for _, tile := range tiles {
		for _, other := range tiles {
			area := int((math.Abs(float64(tile[0]-other[0])) + 1) * (math.Abs(float64(tile[1]-other[1])) + 1))
			if area > hi {
				hi = area
			}
		}
	}

	return hi
}
