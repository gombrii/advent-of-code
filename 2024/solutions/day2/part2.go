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
	fmt.Printf("Checking %v\n", report)
	if safe(report) {
		fmt.Printf("Report %v is safe\n", report)
		return true
	}
	for i := range report {
		if safe(append(report[:i:i], report[i+1:]...)) {
			fmt.Printf("Report %v is safe with %d removed\n", report, report[i])
			return true
		}
	}
	fmt.Printf("Report %v is unsafe\n", report)
	return false
}