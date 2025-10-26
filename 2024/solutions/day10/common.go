package day10

import (
	"strconv"

	"github.com/gombrii/Advent-of-code/shared/dat"
	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/iterate"
	"github.com/gombrii/Advent-of-code/shared/mat"
)

var dirs = []dat.Vec[int]{
	{X: -1, Y: 0},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
}

const goal = "9"

func findTrailheads(matrix [][]string) map[dat.Vec[int]]bool {
	trailheads := make(map[dat.Vec[int]]bool)
	for pos, val := range iterate.Matrix(matrix) {
		if val == "0" {
			trailheads[pos] = true
		}
	}
	return trailheads
}

func paths(matrix [][]string, from dat.Vec[int], ends []dat.Vec[int]) []dat.Vec[int] {
	if matrix[from.Y][from.X] == goal {
		return append(ends, from)
	}

	return turns(matrix, from, ends)
}

func turns(matrix [][]string, from dat.Vec[int], ends []dat.Vec[int]) []dat.Vec[int] {
	for _, dir := range dirs {
		to := from.Add(dir)
		if !mat.InMatrix(matrix, to.X, to.Y) {
			continue
		}
		currHeight, err := strconv.Atoi(matrix[from.Y][from.X])
		exit.If(err)
		nextHeight, err := strconv.Atoi(matrix[to.Y][to.X])
		exit.If(err)

		if nextHeight-currHeight == 1 {
			ends = paths(matrix, to, ends)
		}
	}
	return ends
}
