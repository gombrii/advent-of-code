// Package day2 solves puzzle available on https://adventofcode.com/2025/day/2
package day2

import (
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Parts(data, ",")
	from, to := parseData(in)

	return totalInvalid(from, to, validPart1)
}

func validPart1(id string) bool {
	return len(id)%2 != 0 || id[:len(id)/2] != id[len(id)/2:]
}
