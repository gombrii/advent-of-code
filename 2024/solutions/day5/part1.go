package day5

import (
	"github.com/gombrii/Advent-of-code/shared/input"
	"github.com/gombrii/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day5", "part1", Part1)
}

func Part1(file string) any {
	in := input.Slice(file)

	rules, updates := parse(in)
	correctUpdates := findCorrect(rules, updates)
	sum := sum(correctUpdates)

	return sum
}

func findCorrect(rules map[string][]string, updates []map[string]int) []map[string]int {
	correctUpdates := make([]map[string]int, 0, len(updates)/2)
	for _, update := range updates {
		if correct(rules, update) {
			correctUpdates = append(correctUpdates, update)
		}
	}
	return correctUpdates
}
