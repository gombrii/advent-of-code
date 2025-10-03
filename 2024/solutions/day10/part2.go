package day10

import (
	"github.com/gomsim/Advent-of-code/shared/dat"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day10", "part2", Part2)
}

func Part2(file string) any {
	in := input.Matrix(file)

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
