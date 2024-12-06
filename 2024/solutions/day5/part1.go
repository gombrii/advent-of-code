package day5

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day5", Part1)
}

func Part1(file string) {
	in := input.Slice(file)

	rules, updates := parse(in)
	correctUpdates := findCorrect(rules, updates)
	sum := sum(correctUpdates)

	fmt.Println(sum)
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
