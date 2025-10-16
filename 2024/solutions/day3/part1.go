package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	parse1 "github.com/gombrii/Advent-of-code/shared/parse"
)

var mulPattern = regexp.MustCompile(`mul\(\d+,\d+\)`)
var mulWash = regexp.MustCompile(`mul|\(|\)`)

func Part1(data []byte) any {
	in := parse1.String(data)

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
