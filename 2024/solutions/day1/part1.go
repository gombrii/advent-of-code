package day1

import (
	"math"
	"sort"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day1", "part1", Part1)
}

func Part1(file string) any {
	in := input.Slice(file)

	a, b := intArrays(in)
	sort.Ints(a)
	sort.Ints(b)
	dist := totalDist(a, b)

	return dist
}

func totalDist(a []int, b []int) int {
	acc := 0
	for i := range a {
		acc += dist(a[i], b[i])
	}
	return acc
}

func dist(a int, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}
