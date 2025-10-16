package day10

import (
	"github.com/gombrii/Advent-of-code/shared/dat"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Matrix(data, " ")

	trailheads := findTrailheads(in)
	score := scoreTrails(in, trailheads)

	return score
}

func scoreTrails(matrix [][]string, trailheads map[dat.Vec[int]]bool) int {
	score := 0
	for trailhead := range trailheads {
		score += len(paths(matrix, trailhead, []dat.Vec[int]{}))
	}
	return score
}
