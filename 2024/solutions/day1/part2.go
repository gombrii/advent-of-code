package day1

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)

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
