package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/common/input"
)

var max = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var outcome = regexp.MustCompile(`(\d+) (\w+)`)

func main() {
	in := input.Array()

	sum := sumPossibleGames(in)

	fmt.Println(sum)
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
