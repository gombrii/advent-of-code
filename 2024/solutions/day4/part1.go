package day4

import (
	"fmt"
	"regexp"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/matrices"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

var forw = regexp.MustCompile("XMAS")
var bacw = regexp.MustCompile("SAMX")

func init() {
	registrar.Register("2024", "day4", "part1", Part1)
}

func Part1(file string) any {
	in := input.OldBadMatrix(file)

	hor := in
	ver := ver(in)
	diagNeg := diagNeg(in)
	diagPos := diagPos(in)

	occ := count(hor) + count(ver) + count(diagNeg) + count(diagPos)

	return occ
}

func ver(matrix [][]byte) [][]byte {
	out := blankCopy(matrix)
	for x := range matrix[0] {
		for y, val := range matrices.Ver(matrix, x) {
			out[x][y] = val
		}
		fmt.Printf("ver X start (%d): %s\n", x, string(out[x])) // Debug output
	}
	return out
}

func diagNeg(matrix [][]byte) [][]byte {
	out := make([][]byte, 0, len(matrix)*2)
	for x := range matrix[0] {
		outRow := make([]byte, 0, len(matrix)*2)
		for _, val := range matrices.DiagNegX(matrix, x) {
			outRow = append(outRow, val)
		}
		fmt.Printf("diagNeg X start (%d): %s\n", x, string(outRow)) // Debug output
		out = append(out, outRow)
	}
	for y := 1; y < len(matrix); y++ {
		outRow := make([]byte, 0, len(matrix)*2)
		for _, val := range matrices.DiagNegY(matrix, y) {
			outRow = append(outRow, val)
		}
		fmt.Printf("diagNeg Y start (%d): %s\n", y, string(outRow)) // Debug output
		out = append(out, outRow)
	}

	return out
}

func diagPos(matrix [][]byte) [][]byte {
	out := make([][]byte, 0, len(matrix)*2)
	for x := range matrix[0] {
		outRow := make([]byte, 0, len(matrix)*2)
		for _, val := range matrices.DiagPosX(matrix, x) {
			outRow = append(outRow, val)
		}
		fmt.Printf("diagPos X start (%d): %s\n", x, string(outRow)) // Debug output
		out = append(out, outRow)
	}
	for y := 1; y < len(matrix); y++ {
		outRow := make([]byte, 0, len(matrix)*2)
		for _, val := range matrices.DiagPosY(matrix, y) {
			outRow = append(outRow, val)
		}
		fmt.Printf("diagPos Y start (%d): %s\n", y, string(outRow)) // Debug output
		out = append(out, outRow)
	}

	return out
}

func blankCopy(matrix [][]byte) [][]byte {
	blank := make([][]byte, len(matrix))
	for i := range blank {
		blank[i] = make([]byte, len(matrix[0]))
	}
	return blank
}

func count(rows [][]byte) int {
	acc := 0
	for _, row := range rows {
		forwardMatches := len(forw.FindAll(row, -1))
		backwardMatches := len(bacw.FindAll(row, -1))
		fmt.Printf("Row: %s, Forward: %d, Backward: %d\n", string(row), forwardMatches, backwardMatches) // Debug output
		acc += forwardMatches
		acc += backwardMatches
	}
	return acc
}
