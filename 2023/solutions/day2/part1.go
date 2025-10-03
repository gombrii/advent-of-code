package day2

import (
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2023", "day2", "part1", Part1)
}

var max = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Part1(file string) any {
	in := input.Slice(file)

	sum := sumPossibleGames(in)

	return sum
}

func sumPossibleGames(games []string) int {
	acc := 0
	for i, game := range games {
		if possibleGame(game) {
			acc += gameID(i)
		}
	}
	return acc
}

func possibleGame(game string) bool {
	for _, round := range strings.Split(game, ";") {
		if !possibleRound(round) {
			return false
		}
	}
	return true
}

func possibleRound(round string) bool {
	const count, colour = 1, 2
	curr := make(map[string]int)

	for _, group := range outcome.FindAllStringSubmatch(round, -1) {
		num, _ := strconv.Atoi(group[count])
		curr[group[colour]] = num
	}

	for key, val := range curr {
		if val > max[key] {
			return false
		}
	}

	return true
}

func gameID(row int) int {
	return row + 1
}
