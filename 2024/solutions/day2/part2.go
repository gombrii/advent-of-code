package day2

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part2("2024", "day2", Part2)
}

func Part2(file string) {
	in := input.Array(file)
	washed := wash(in)
	numSafe := numSafe(washed, safeWithDampening)

	fmt.Println(numSafe)
}

func safeWithDampening(report []int) bool {
	tree := createTree(report)
	for _, variant := range tree {
		if safe(variant) {
			return true
		}
	}
	return false
}

func createTree(report []int) [][]int {
	tree := make([][]int, len(report)+1)
	tree[0] = report
	for i := range report {
		tree[i+1] = removeItem(report, i)
	}
	return tree
}

func removeItem(arr []int, index int) []int {
	new := make([]int, len(arr)-1)
	copy(new[:index], arr[:index])
	copy(new[index:], arr[index+1:])
	return new
}
