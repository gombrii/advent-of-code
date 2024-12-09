package mat

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/data"
)

func InBounds[T any](matrix [][]T, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix[0]) && y < len(matrix)
}

func PrintMatrix[T any](matrix [][]T) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			fmt.Print(matrix[y][x])
		}
		fmt.Println()
	}
}

func PrintMatrixReplace[T any](matrix [][]T, replace map[data.Vec[int]]T) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if val, exists := replace[data.Vec[int]{X: x, Y: y}]; exists {
				fmt.Print(val)
			} else {
				fmt.Print(matrix[y][x])
			}

		}
		fmt.Println()
	}
}
