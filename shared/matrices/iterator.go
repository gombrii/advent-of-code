package matrices

import (
	"iter"
)

func Ver(matrix [][]byte, x int) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for y, row := range matrix {
			if !yield(y, row[x]) {
				return
			}
		}
	}
}

func DiagNegX(matrix [][]byte, X int) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for y := 0; y+X < len(matrix[0]); y++ {
			if !yield(y, matrix[y][y+X]) {
				return
			}
		}
	}
}

func DiagNegY(matrix [][]byte, Y int) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for y, row := range matrix[Y:] {
			if !yield(y, row[y]) {
				return
			}
		}
	}
}

func DiagPosX(matrix [][]byte, X int) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for y := 0; y <= X; y++ {
			if !yield(y, matrix[y][X-y]) {
				return
			}
		}
	}
}

func DiagPosY(matrix [][]byte, Y int) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for y := 0; y < len(matrix)-Y; y++ {
			if !yield(y, matrix[y+Y][len(matrix[0])-y-1]) {
				return
			}
		}
	}
}
