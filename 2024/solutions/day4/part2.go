package day4

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/matrices"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part2("2024", "day4", Part2)
}

func Part2(file string) {
	in := input.OldBadMatrix(file)

	occ := countX(in)

	fmt.Println(occ)
}

func countX(matrix [][]byte) int {
	acc := 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			if xmas(view(matrix, x, y)) {
				acc++
			}
		}
	}
	return acc
}

func xmas(view [][]byte) bool {
	neg, pos := make([]byte, 3), make([]byte, 3)
	for i, val := range matrices.DiagNegY(view, 0) {
		neg[i] = val
	}
	for i, val := range matrices.DiagPosY(view, 0) {
		pos[i] = val
	}

	return (string(neg) == "MAS" || string(neg) == "SAM") && (string(pos) == "MAS" || string(pos) == "SAM")
}

func view(matrix [][]byte, x int, y int) (view [][]byte) {
	view = make([][]byte, 3)
	for i, row := range matrix[y-1 : y+2] {
		view[i] = row[x-1 : x+2]
	}
	return view
}
