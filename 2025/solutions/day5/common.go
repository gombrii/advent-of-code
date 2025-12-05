// Package day5 solves puzzle available on https://adventofcode.com/2025/day/5
package day5

import (
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
)

func compile(data []string) (ranges map[int]int, ids []int) {
	ranges = map[int]int{}
	ids = []int{}

	delimiter := 0
	for i, line := range data {
		delimiter = i
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		beg, err := strconv.Atoi(parts[0])
		exit.If(err)
		end, err := strconv.Atoi(parts[1])
		exit.If(err)

		merge(beg, end, ranges)
	}

	for _, line := range data[delimiter+1:] {
		id, err := strconv.Atoi(line)
		exit.If(err)
		ids = append(ids, id)
	}

	return ranges, ids
}

func merge(beg, end int, ranges map[int]int) {
	for b, e := range ranges {
		switch {
		case beg >= b && end <= e:
			return
		case beg >= b && beg <= e && end > e:
			beg = b
		case beg < b && end >= b && end <= e:
			delete(ranges, b)
			end = e
		case beg <= b && end >= e:
			delete(ranges, b)
		}
	}

	ranges[beg] = end
}
