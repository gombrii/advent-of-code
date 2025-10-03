package day1

import (
	"regexp"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2023", "day1", "part2", Part2)
}

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

func Part2(file string) any {
	in := input.Slice(file)

	formatted := format(in)
	numStrings := findAll(formatted)
	sum := parseAndSum(numStrings)

	return sum
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
