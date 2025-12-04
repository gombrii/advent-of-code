// Package day4 solves puzzle available on https://adventofcode.com/2025/day/4
package day4

import (
	"iter"
	"math"

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

func accessible(matrix [][]string, x, y int) bool {
	if matrix[y][x] != "@" {
		return false
	}

	acc := 0
	for x, y := range Neighbours(matrix, x, y) {
		if matrix[y][x] == "@" {
			acc++
		}
	}

	return acc < 4
}

func Neighbours(matrix [][]string, x, y int) iter.Seq2[int, int] {
	minX := int(math.Max(float64(x-1), 0))
	minY := int(math.Max(float64(y-1), 0))
	maxX := int(math.Min(float64(x+2), float64(len(matrix[0]))))
	maxY := int(math.Min(float64(y+2), float64(len(matrix))))

	return func(yield func(int, int) bool) {
		for yi := minY; yi < maxY; yi++ {
			for xi := minX; xi < maxX; xi++ {
				if (xi != x || yi != y) && !yield(xi, yi) {
					return
				}
			}
		}
	}
}
