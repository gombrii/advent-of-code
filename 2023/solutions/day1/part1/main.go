package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gomsim/Advent-of-code/common/exit"
	"github.com/gomsim/Advent-of-code/common/input"
)

var pattern = regexp.MustCompile(`\d`)

func main() {
	in := input.Array()

	numStrings := findAll(in)
	sum := parseAndSum(numStrings)

	fmt.Println(sum)
}

func firstLast(line string) (first string, last string) {
	matches := pattern.FindAllString(line, -1)
	return matches[0], matches[len(matches)-1]
}

func findAll(lines []string) [][]string {
	numbers := make([][]string, len(lines))
	for i, line := range lines {
		first, last := firstLast(line)
		numbers[i] = []string{first, last}
	}
	return numbers
}

func parseAndSum(strings [][]string) int {
	acc := 0
	for _, str := range strings {
		num, err := strconv.Atoi(fmt.Sprint(str[0], str[1]))
		if err != nil {
			exit.Errorf("couldn't parse number: %v", err)
		}
		acc += num
	}
	return acc
}
