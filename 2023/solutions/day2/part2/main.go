package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/common/input"
)

var outcome = regexp.MustCompile(`(\d+) (\w+)`)

func main() {
	in := input.Array()

	sum := sumPower(in)

	fmt.Println(sum)
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
