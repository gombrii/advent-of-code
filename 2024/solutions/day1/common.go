package day1

import (
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
)

func intArrays(data []string) (a []int, b []int) {
	for _, line := range data {
		tokens := strings.Split(line, "   ")
		ai, err := strconv.Atoi(tokens[0])
		exit.If(err)
		bi, err := strconv.Atoi(tokens[1])
		exit.If(err)
		a = append(a, ai)
		b = append(b, bi)
	}
	return a, b
}
