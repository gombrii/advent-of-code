package day10

import (
	"fmt"
	"strconv"

	"github.com/gomsim/Advent-of-code/shared/data"
	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/mat"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day10", "part1", Part1)
}

var dirs = []data.Vec[int]{
	{X: -1, Y: 0},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
}

const goal = "9"

func Part1(file string) {
	in := input.Matrix(file)

	trailheads := findTrailheads(in)
	score := trails(in, trailheads)

	fmt.Println(score)
}

func findTrailheads(matrix [][]string) map[data.Vec[int]]bool {
	trailheads := make(map[data.Vec[int]]bool)
	for pos, val := range data.Miter(matrix) {
		if val == "0" {
			trailheads[pos] = true
		}
	}
	return trailheads
}

func trails(matrix [][]string, trailheads map[data.Vec[int]]bool) int {
	score := 0
	for trailhead := range trailheads {
		score += len(trail(matrix, trailhead, make(map[data.Vec[int]]bool)))
	}
	return score
}

func trail(matrix [][]string, from data.Vec[int], ends map[data.Vec[int]]bool) map[data.Vec[int]]bool {
	if matrix[from.Y][from.X] == goal {
		ends[from] = true
		return ends
	}

	//fmt.Println()
	//time.Sleep(time.Millisecond * 100)
	//mat.PrintMatrixReplace(matrix, mat.Replace[string]{
	//	"█": map[data.Vec[int]]bool{
	//		{X: from.X, Y: from.Y}: true,
	//	},
	//	" ": ends,
	//})

	return turns(matrix, from, ends)
}

func turns(matrix [][]string, from data.Vec[int], ends map[data.Vec[int]]bool) map[data.Vec[int]]bool {
	for _, dir := range dirs {
		nextLoc := from.Add(dir)
		if !mat.InBounds(matrix, nextLoc.X, nextLoc.Y) {
			continue
		}
		currHeight, err := strconv.Atoi(matrix[from.Y][from.X])
		exit.If(err)
		nextHeight, err := strconv.Atoi(matrix[nextLoc.Y][nextLoc.X])
		exit.If(err)

		if nextHeight-currHeight == 1 {
			ends = trail(matrix, nextLoc, ends)
		}
	}

	return ends
}
