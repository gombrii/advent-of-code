package day1

import (
	"math"
	"sort"

	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)

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
