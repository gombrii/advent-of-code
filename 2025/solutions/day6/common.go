// Package day6 solves puzzle available on https://adventofcode.com/2025/day/6
package day6

func compile(data []string) ([][]string, []rune) {
	rows := make([][]string, 0)
	ops := make([]rune, 0)

	delimiters := make([]int, 0)
	for i, r := range data[len(data)-1] {
		if r == '+' || r == '*' {
			delimiters = append(delimiters, i)
			ops = append(ops, r)
		}
	}
	for _, line := range data[:len(data)-1] {
		row := make([]string, 0)
		for i := range delimiters {
			end := len(line)
			if i+1 < len(delimiters) {
				end = delimiters[i+1] - 1
			}
			row = append(row, line[delimiters[i]:end])
		}
		rows = append(rows, row)
	}

	return rows, ops
}

func operate(a, b int, op rune) int {
	if op == '*' {
		if a == 0 {
			a = 1
		}
		return a * b
	}

	return a + b
}
