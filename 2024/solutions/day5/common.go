package day5

import (
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
)

func parse(input []string) (rules map[string][]string, updates []map[string]int) {
	rules = make(map[string][]string, len(input))
	i := 0
	for input[i] != "" {
		tokens := strings.Split(input[i], "|")
		rules[tokens[0]] = append(rules[tokens[0]], tokens[1])
		i++
	}

	i++
	for i < len(input) {
		update := make(map[string]int, len(input[i])*2)
		for i, page := range strings.Split(input[i], ",") {
			update[page] = i
		}
		updates = append(updates, update)
		i++
	}
	return rules, updates
}

func correct(rules map[string][]string, update map[string]int) bool {
	for val, i := range update {
		for _, mustAfter := range rules[val] {
			if aftPos, present := update[mustAfter]; present && aftPos < i {
				return false
			}
		}
	}
	return true
}

func sum(updates []map[string]int) int {
	acc := 0

	for _, update := range updates {
		for page, i := range update {
			if i == len(update)/2 {
				num, err := strconv.Atoi(page)
				exit.If(err)
				acc += num
				break
			}
		}
	}

	return acc
}
