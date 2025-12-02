// Package day2 solves puzzle available on https://adventofcode.com/2025/day/2
package day2

import (
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Parts(data, ",")
	from, to := parseData(in)

	invalid := findInvalidTotal(from, to)

	acc := 0
	for _, id := range invalid {
		acc += id
	}

	return acc
}

func parseData(data []string) ([]int, []int) {
	from := make([]int, len(data))
	to := make([]int, len(data))

	for i, l := range data {
		parts := strings.Split(l, "-")

		f, err := strconv.Atoi(parts[0])
		exit.If(err)
		t, err := strconv.Atoi(parts[1])
		exit.If(err)

		from[i] = f
		to[i] = t
	}

	return from, to
}

func findInvalidTotal(from, to []int) []int {
	invalidIDs := make([]int, 0)

	for i := range len(from) {
		invalidIDs = append(invalidIDs, findInvalidInRange(from[i], to[i])...)
	}

	return invalidIDs
}

func findInvalidInRange(from, to int) []int {
	invalidIDs := make([]int, to-from)

	for id := from; id <= to; id++ {
		idStr := strconv.Itoa(id)
		if invalid(idStr) {
			invalidIDs = append(invalidIDs, id)
		}
	}

	return invalidIDs
}

func invalid(id string) bool {
	return len(id)%2 == 0 && id[:len(id)/2] == id[len(id)/2:]
}
