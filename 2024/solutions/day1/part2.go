package day1

import (
	"github.com/gombrii/Advent-of-code/shared/input"
	"github.com/gombrii/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day1", "part2", Part2)
}

func Part2(file string) any {
	in := input.Slice(file)

	a, b := intArrays(in)
	occurances := occurances(b)
	score := simscore(a, occurances)

	return score
}

func occurances(numbers []int) map[int]int {
	occ := make(map[int]int)
	for _, num := range numbers {
		occ[num] = occ[num] + 1
	}
	return occ
}

func simscore(numbers []int, occ map[int]int) int {
	acc := 0
	for _, num := range numbers {
		acc += occ[num] * num
	}
	return acc
}
