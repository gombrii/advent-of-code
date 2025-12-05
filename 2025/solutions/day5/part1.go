// Package day5 solves puzzle available on https://adventofcode.com/2025/day/5
package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

// OLD GUESS
func Part1(data []byte) any {
	in := parse.Lines(data)
	ranges, ids := compile(in)

	fmt.Println("COMPILED RANGES", ranges)

	acc := 0
	for _, id := range ids {
		if fresh(id, ranges) {
			acc++
		}
	}

	return acc
}

func fresh(id int, ranges map[int]int) bool {
	for b, e := range ranges {
		if id >= b && id <= e {
			return true
		}
	}

	return false
}

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

		if !merge(beg, end, ranges) {
			ranges[beg] = end
		}
	}

	for _, line := range data[delimiter+1:] {
		id, err := strconv.Atoi(line)
		exit.If(err)
		ids = append(ids, id)
	}

	return ranges, ids
}

func merge(beg, end int, ranges map[int]int) bool {
	fmt.Println("BUILDING RANGES", ranges)
	merged := false
	for b, e := range ranges {
		switch {
		case beg >= b && beg <= e && end >= b && end <= e:
			fmt.Println("I am b", beg, "e", end)
			fmt.Println("For b", b, "e", e, "Do nothing")
			// Do nothing
			return true
		case beg >= b && beg <= e && end > e:
			fmt.Println("I am b", beg, "e", end)
			fmt.Println("For b", b, "e", e, "Update end")
			// Update end
			merged = true
			ranges[b] = end
			beg = b
		case beg < b && end >= b && end <= e:
			fmt.Println("I am b", beg, "e", end)
			fmt.Println("For b", b, "e", e, "Update beginning")
			//Update beginning
			merged = true
			delete(ranges, b)
			ranges[beg] = e
		case beg <= b && end >= e:
			fmt.Println("I am b", beg, "e", end)
			fmt.Println("For b", b, "e", e, "Consume old")
			//Delete old, replace with new
			merged = true
			delete(ranges, b)
			ranges[beg] = end
		}
	}
	return merged
}
