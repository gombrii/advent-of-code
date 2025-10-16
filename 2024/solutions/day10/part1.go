package day10

import (
	"github.com/gombrii/Advent-of-code/shared/dat"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Matrix(data, " ")

	trailheads := findTrailheads(in)
	score := scoreTrailheads(in, trailheads)

	return score
}

func scoreTrailheads(matrix [][]string, trailheads map[dat.Vec[int]]bool) int {
	score := 0
	for trailhead := range trailheads {
		trails := paths(matrix, trailhead, []dat.Vec[int]{})
		distinct := make(map[dat.Vec[int]]bool)
		for _, trail := range trails {
			distinct[trail] = true
		}
		score += len(distinct)
	}
	return score
}
