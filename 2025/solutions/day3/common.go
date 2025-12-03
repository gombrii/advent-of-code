// Package day3 solves puzzle available on https://adventofcode.com/2025/day/3
package day3

import (
	"strconv"

	"github.com/gombrii/Advent-of-code/shared/exit"
)

func joltage(bank string, size int) int {
	res := ""
	for n := -size + 1; n <= 0; n++ {
		var hJoltage byte = '0'
		hPos := 0
		for i := len(bank) + n - 1; i >= 0; i-- {
			if bank[i] == '*' {
				break
			}

			if bank[i] >= hJoltage {
				hJoltage = bank[i]
				hPos = i
			}
		}
		bank = bank[:hPos] + "*" + bank[hPos+1:]
		res += string(hJoltage)
	}
	out, err := strconv.Atoi(res)
	exit.If(err)

	return out
}
