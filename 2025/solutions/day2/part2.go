// Package day2 solves puzzle available on https://adventofcode.com/2025/day/2
package day2

import (
	"strings"

	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Parts(data, ",")
	from, to := parseData(in)

	return totalInvalid(from, to, validPart2)
}

func validPart2(id string) bool {
	for i := range len(id) {
		seq := id[:i]
		if strings.Count(id, seq)*len(seq) == len(id) {
			return false
		}
	}

	return true
}
