package render

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/dat"
)

func Matrix[T any](matrix [][]T) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			fmt.Print(matrix[y][x])
		}
		fmt.Println()
	}
}

type Replace[T comparable] map[T]map[dat.Vec[int]]bool

func MatrixReplace[T comparable](matrix [][]T, replace Replace[T]) {
	replacements := make(map[dat.Vec[int]]T)
	for rep, coordinates := range replace {
		for coordinate := range coordinates {
			replacements[coordinate] = rep
		}
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if val, exists := replacements[dat.Vec[int]{X: x, Y: y}]; exists {
				fmt.Print(val)
			} else {
				fmt.Print(matrix[y][x])
			}

		}
		fmt.Println()
	}
}
