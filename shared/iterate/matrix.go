package iterate

import (
	"iter"

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
