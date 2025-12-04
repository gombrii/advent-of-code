package iterate

import (
	"iter"
	"math"

	"github.com/gombrii/Advent-of-code/shared/dat"
)

func Matrix(matrix [][]string) iter.Seq2[dat.Vec[int], string] {
	return func(yield func(dat.Vec[int], string) bool) {
		for y := 0; y < len(matrix); y++ {
			for x := 0; x < len(matrix[0]); x++ {
				if !yield(dat.Vec[int]{X: x, Y: y}, matrix[y][x]) {
					return
				}
			}
		}
	}
}

func Neighbours(x, y int, matrix [][]string) iter.Seq2[int, int] {
	minX := int(math.Max(float64(x-1), 0))
	minY := int(math.Max(float64(y-1), 0))
	maxX := int(math.Min(float64(x+2), float64(len(matrix))))
	maxY := int(math.Min(float64(y+2), float64(len(matrix[0]))))

	return func(yield func(int, int) bool) {
		for yi, row := range matrix[minY:maxY] {
			for xi := range row[minX:maxX] {
				if xi != x && yi != y && !yield(x, y) {
					return
				}
			}
		}
	}
}
