// Package day6 solves puzzle available on https://adventofcode.com/2025/day/6
package day6

import (
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part2(data []byte) any {
	in := parse.Lines(data)
	rows, ops := compile(in)

	acc := 0
	for x := range len(rows[0]) {
		acc += calculateLikeACephalopod(x, rows, ops)
	}

	return acc
}

func calculateLikeACephalopod(col int, rows [][]string, ops []rune) int {
	acc := 0

	for x := range len(rows[0][col]) {
		s := ""
		for y := range rows {
			s += string(rows[y][col][x])
		}
		num, err := strconv.Atoi(strings.TrimSpace(s))
		exit.If(err)

		acc = operate(acc, num, ops[col])
	}

	return acc
}
