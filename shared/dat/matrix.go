package dat

import "iter"

func Miter(matrix [][]string) iter.Seq2[Vec[int], string] {
	return func(yield func(Vec[int], string) bool) {
		for y := 0; y < len(matrix); y++ {
			for x := 0; x < len(matrix[0]); x++ {
				if !yield(Vec[int]{X: x, Y: y}, matrix[y][x]) {
					return
				}
			}
		}
	}
}
