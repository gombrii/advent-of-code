package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gomsim/Advent-of-code/common/exit"
)

const all = -1

var pattern = regexp.MustCompile("[0-9]")

func firstLast(line string) (first string, last string) {
	matches := pattern.FindAllString(line, all)
	return matches[0], matches[len(matches)-1]
}

func FindAll(lines []string) [][]string {
	numbers := make([][]string, len(lines))
	for i, line := range lines {
		first, last := firstLast(line)
		numbers[i] = []string{first, last}
	}
	return numbers
}

func Parse(strings [][]string) []int {
	numbers := make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(fmt.Sprint(str[0], str[1]))
		if err != nil {
			exit.Errorf("couldn't parse number: %v", err)
		}
		numbers[i] = num
	}
	return numbers
}

func Sum(numbers []int) int {
	acc := 0
	for _, number := range numbers {
		acc += number
	}
	return acc
}
