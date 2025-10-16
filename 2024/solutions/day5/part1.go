package day5

import (
	parse1 "github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse1.Lines(data)

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
