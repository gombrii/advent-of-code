package day6

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day6", Part1)
}

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

func Part1(file string) {
	matrix := input.Matrix(file)

	guard := findGuard(matrix)
	guard = passTime(matrix, guard)

	fmt.Println(len(guard.trace))
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
		//printMatrix(matrix) // debug
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
