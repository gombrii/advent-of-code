// Package day6 solves puzzle available on https://adventofcode.com/2025/day/6
package day6

import (
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse.Lines(data)
	rows, ops := compile(in)

	acc := 0
	for x := range len(rows[0]) {
		acc += calculateLikeAHuman(x, rows, ops)
	}

	return acc
}

func calculateLikeAHuman(col int, rows [][]string, ops []rune) int {
	acc := 0
	for y := range rows {
		num, err := strconv.Atoi(strings.TrimSpace(rows[y][col]))
		exit.If(err)
		
		acc = operate(acc, num, ops[col])
	}

	return acc
}
