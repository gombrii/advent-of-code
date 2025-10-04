package day6

import (
	"github.com/gombrii/Advent-of-code/shared/input"
	"github.com/gombrii/Advent-of-code/shared/registry"
	"github.com/gombrii/Advent-of-code/shared/render"
)

func init() {
	registry.Register("2024", "day6", "part2", Part2)
}

func Part2(file string) any {
	matrix := input.Matrix(file)
	render.Matrix(matrix)

	return "NOT IMPLEMENTED!"
}
