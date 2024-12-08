package day8

import (
	"regexp"

	"github.com/gomsim/Advent-of-code/shared/data"
)

var antenna = regexp.MustCompile(`\w|\d`)

func parse(matrix [][]string) catalogue {
	catal := make(catalogue)

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			marking := matrix[y][x]
			if antenna.MatchString(marking) {
				catal[marking] = append(catal[marking], data.Vec[int]{X: x, Y: y})
			}
		}
	}

	return catal
}
