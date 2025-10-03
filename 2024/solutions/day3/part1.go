package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day3", "part1", Part1)
}

var mulPattern = regexp.MustCompile(`mul\(\d+,\d+\)`)
var mulWash = regexp.MustCompile(`mul|\(|\)`)

func Part1(file string) any {
	in := input.String(file)

	result := runProgram(in)

	return result
}

func runProgram(input string) int {
	acc := 0
	for _, match := range mulPattern.FindAllString(input, -1) {
		a, b := parse(match)
		acc += perform(a, b)
	}
	return acc
}

func parse(input string) (int, int) {
	washed := mulWash.ReplaceAllString(input, "")
	tokens := strings.Split(washed, ",")
	a, err := strconv.Atoi(tokens[0])
	exit.If(err)
	b, err := strconv.Atoi(tokens[1])
	exit.If(err)
	return a, b
}

func perform(a int, b int) int {
	return a * b
}
