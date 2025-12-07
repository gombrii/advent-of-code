// Package day7 solves puzzle available on https://adventofcode.com/2025/day/7
package day7

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

type node struct {
	x int
	y int
}

func Part2(data []byte) any {
	in := parse.Matrix(data)

	return compile(node{startX(in), 0}, in, map[node]int{})
}

func compile(n node, matrix [][]string, cache map[node]int) int {
	v, ok := cache[n]
	if !ok {
		for y, row := range matrix[n.y:] {
			if row[n.x] == "^" {
				left := node{x: n.x - 1, y: n.y + y}
				right := node{x: n.x + 1, y: n.y + y}
				cache[left] = compile(left, matrix, cache)
				cache[right] = compile(right, matrix, cache)
				return cache[left] + cache[right]
			}
		}

		return 1
	}

	return v
}
