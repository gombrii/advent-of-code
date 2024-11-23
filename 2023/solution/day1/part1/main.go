package main

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/2023/solution/day1"
	"github.com/gomsim/Advent-of-code/common/input"
)

func main() {
	numStrings := day1.FindAll(input.Array())
	numbers := day1.Parse(numStrings)
	sum := day1.Sum(numbers)

	fmt.Println(sum)
}
