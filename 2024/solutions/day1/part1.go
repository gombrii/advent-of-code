package day1

import (
	"fmt"
	"math"
	"sort"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day1", "part1", Part1)
}

func Part1(file string) {
	in := input.Slice(file)

	a, b := intArrays(in)
	sort.Ints(a)
	sort.Ints(b)
	dist := totalDist(a, b)

	fmt.Println(dist)
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
