package day6

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

const wall = "#"

var (
	west  = vec{-1, 0}
	north = vec{0, -1}
	east  = vec{1, 0}
	south = vec{0, 1}
)

type vec struct {
	x int
	y int
}

type guard struct {
	pos   vec
	dir   vec
	trace map[vec]bool
}

func Part1(data []byte) any {
	matrix := parse.Matrix(data, "")

	guard := findGuard(matrix)
	guard = passTime(matrix, guard)

	return len(guard.trace)
}

func findGuard(matrix [][]string) guard {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == "^" {
				return guard{
					pos:   vec{x, y},
					dir:   north,
					trace: map[vec]bool{},
				}
			}
		}
	}
	panic("Guard not found")
}

func passTime(matrix [][]string, guard guard) guard {
	for guard.pos.x+guard.dir.x <= len(matrix[0]) && guard.pos.y+guard.dir.y <= len(matrix) {
		// print(matrix, guard) // debug
		guard = tick(matrix, guard)
	}
	return guard
}

func tick(matrix [][]string, guard guard) guard {
	guard.trace[guard.pos] = true
	if guard.pos.x+guard.dir.x < len(matrix[0]) && guard.pos.y+guard.dir.y < len(matrix) && matrix[guard.pos.y+guard.dir.y][guard.pos.x+guard.dir.x] == wall {
		switch guard.dir {
		case north:
			guard.dir = east
		case east:
			guard.dir = south
		case south:
			guard.dir = west
		case west:
			guard.dir = north
		}
	}
	guard.pos.x = guard.pos.x + guard.dir.x
	guard.pos.y = guard.pos.y + guard.dir.y
	return guard
}

//func print(matrix [][]string, guard guard) {
//	rocks := make([]render.G, 0)
//	for pos, c := range iterate.Matrix(matrix) {
//		if c == "#" {
//			rocks = append(rocks, render.G{Symbol: '#', X: pos.X, Y: pos.Y})
//		}
//	}
//
//	render.Animate(250, render.GString(len(matrix[0]), len(matrix), '.',
//		append(rocks, render.G{Symbol: '^', X: guard.pos.x, Y: guard.pos.y}),
//	))
//}
