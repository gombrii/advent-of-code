package day10

import (
	"strconv"

	"github.com/gomsim/Advent-of-code/shared/dat"
	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/mat"
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
	for pos, val := range dat.Miter(matrix) {
		if val == "0" {
			trailheads[pos] = true
		}
	}
	return trailheads
}

func trail(matrix [][]string, from dat.Vec[int], ends map[dat.Vec[int]]bool) (map[dat.Vec[int]]bool, int) {
	//fmt.Println()
	//time.Sleep(time.Millisecond * 100)
	//mat.PrintMatrixReplace(matrix, mat.Replace[string]{
	//	" ": ends,
	//	"█": map[data.Vec[int]]bool{
	//		{X: from.X, Y: from.Y}: true,
	//	},
	//})

	if matrix[from.Y][from.X] == goal {
		ends[from] = true
		return ends, 1
	}

	goals, rating := turns(matrix, from, ends)

	//fmt.Println()
	//time.Sleep(time.Millisecond * 100)
	//mat.PrintMatrixReplace(matrix, mat.Replace[string]{
	//	" ": ends,
	//	"█": map[data.Vec[int]]bool{
	//		{X: from.X, Y: from.Y}: true,
	//	},
	//})

	return goals, rating
}

func turns(matrix [][]string, from dat.Vec[int], ends map[dat.Vec[int]]bool) (map[dat.Vec[int]]bool, int) {
	acc := 0
	for _, dir := range dirs {
		to := from.Add(dir)
		if !mat.InBounds(matrix, to.X, to.Y) {
			continue
		}
		currHeight, err := strconv.Atoi(matrix[from.Y][from.X])
		exit.If(err)
		nextHeight, err := strconv.Atoi(matrix[to.Y][to.X])
		exit.If(err)

		if nextHeight-currHeight == 1 {
			newEnds, rating := trail(matrix, to, ends)
			ends = newEnds
			acc += rating
		}
	}
	return ends, acc
}
