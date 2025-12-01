// Package day1 solves puzzle available on https://adventofcode.com/2025/day/1
package day1

import (
	"strconv"

	"github.com/gombrii/Advent-of-code/shared/exit"
)

func parseData(data []string) []int {
	out := make([]int, len(data))
	for i, l := range data {
		n, err := strconv.Atoi(l[1:])
		exit.If(err)

		op := string(l[0])

		out[i] = n
		if op == "L" {
			out[i] *= -1
		}
	}
	return out
}
