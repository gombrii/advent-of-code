package day2

import (
	"math"
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
)

const maxDiff = 3

func wash(input []string) [][]int {
	washedReports := make([][]int, len(input))
	for i, report := range input {
		strNums := strings.Split(report, " ")
		washedReport := make([]int, len(strNums))
		for i, val := range strNums {
			washed, err := strconv.Atoi(val)
			exit.If(err)
			washedReport[i] = washed
		}
		washedReports[i] = washedReport
	}
	return washedReports
}

func safe(report []int) bool {
	ascending := report[0] < report[1]
	for i := 1; i < len(report); i++ {
		if !safeLevels(ascending, report[i-1], report[i]) {
			return false
		}
	}
	return true
}

func numSafe(reports [][]int, safe func([]int) bool) int {
	acc := 0
	for _, report := range reports {
		if safe(report) {
			acc++
		}
	}
	return acc
}

func safeLevels(ascending bool, a int, b int) bool {
	if a == b {
		return false
	}
	if ascending != (a < b) {
		return false
	}
	if math.Abs(float64(a)-float64(b)) > maxDiff {
		return false
	}
	return true
}
