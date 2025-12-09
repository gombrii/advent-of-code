// Package day9 solves puzzle available on https://adventofcode.com/2025/day/9
package day9

import (
	"math"
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
)

const a, b, x, y = 0, 1, 0, 1

func parseData(data []string) [][2]int {
	tiles := make([][2]int, len(data))

	for i, line := range data {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[x])
		exit.If(err)
		y, err := strconv.Atoi(parts[y])
		exit.If(err)

		tiles[i] = [2]int{x, y}
	}

	return tiles
}

func area(a, b [2]int) int {
	return int((math.Abs(float64(a[x]-b[x])) + 1) * (math.Abs(float64(a[y]-b[y])) + 1))
}
