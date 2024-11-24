package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gomsim/Advent-of-code/common/exit"
	"github.com/gomsim/Advent-of-code/common/input"
)

var pattern = regexp.MustCompile(`\d`)

var (
	one   = regexp.MustCompile("one")
	two   = regexp.MustCompile("two")
	three = regexp.MustCompile("three")
	four  = regexp.MustCompile("four")
	five  = regexp.MustCompile("five")
	six   = regexp.MustCompile("six")
	seven = regexp.MustCompile("seven")
	eight = regexp.MustCompile("eight")
	nine  = regexp.MustCompile("nine")
)

func main() {
	in := input.Array()

	formatted := format(in)
	numStrings := findAll(formatted)
	sum := parseAndSum(numStrings)

	fmt.Println(sum)
}

func format(strings []string) []string {
	for i, str := range strings {
		str = one.ReplaceAllString(str, "o1e")
		str = two.ReplaceAllString(str, "t2o")
		str = three.ReplaceAllString(str, "t3e")
		str = four.ReplaceAllString(str, "f4r")
		str = five.ReplaceAllString(str, "f5e")
		str = six.ReplaceAllString(str, "s6x")
		str = seven.ReplaceAllString(str, "s7n")
		str = eight.ReplaceAllString(str, "e8t")
		str = nine.ReplaceAllString(str, "n9e")
		strings[i] = str
	}
	return strings
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
