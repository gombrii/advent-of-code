package day2

import (
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)

	sum := sumPower(in)

	return sum
}

func sumPower(games []string) int {
	acc := 0
	for _, game := range games {
		acc += powerOfGame(game)
	}
	return acc
}

func powerOfGame(game string) int {
	min := make(map[string]int)
	for _, round := range strings.Split(game, ";") {
		for colour, amount := range set(round) {
			if amount > min[colour] {
				min[colour] = amount
			}
		}
	}
	return min["red"] * min["green"] * min["blue"]
}

func set(round string) map[string]int {
	const count, colour = 1, 2
	curr := make(map[string]int)

	for _, group := range outcome.FindAllStringSubmatch(round, -1) {
		num, _ := strconv.Atoi(group[count])
		curr[group[colour]] = num
	}

	return curr
}
