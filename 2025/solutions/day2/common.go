// Package day2 solves puzzle available on https://adventofcode.com/2025/day/2
package day2

import (
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
)

type validator func(string) bool

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

func totalInvalid(from, to []int, validator validator) int {
	var invalidIDs []int
	for i := range len(from) {
		invalidIDs = append(invalidIDs, findInvalidInRange(from[i], to[i], validator)...)
	}

	acc := 0
	for _, id := range invalidIDs {
		acc += id
	}

	return acc
}

func findInvalidInRange(from, to int, valid validator) []int {
	invalidIDs := make([]int, to-from)

	for id := from; id <= to; id++ {
		idStr := strconv.Itoa(id)
		if !valid(idStr) {
			invalidIDs = append(invalidIDs, id)
		}
	}

	return invalidIDs
}
