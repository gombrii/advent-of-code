package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day7", Part1)
}

func Part1(file string) {
	in := input.Slice(file)

	testVals, numbers := extractAll(in)
	total := total(testVals, numbers)

	fmt.Println(total)
}

func total(testVals []int, numbers [][]int) int {
	acc := 0

	for i := range testVals {

		val, possible := possible(testVals[i], numbers[i], nil)
		if possible {
			acc += val
		}
	}

	return acc
}

func possible(testVal int, numbers []int, operators []string) (int, bool) {
	res := calculate(numbers, operators)
	//fmt.Printf("%d=%v %v -> %d\n", testVal, numbers, operators, res)
	if res > testVal {
		return 0, false
	}
	if res == testVal {
		//fmt.Println("Found a winner!")
		return res, true
	}
	if len(operators) >= len(numbers) -1 {
		return 0, false
	}

	withMult := append(operators[:len(operators):len(operators)], "*")
	val, poss := possible(testVal, numbers, withMult)
	if poss {
		return val, true
	}
	withAdd := append(operators[:len(operators):len(operators)], "+")
	return possible(testVal, numbers, withAdd)
}

func calculate(numbers []int, operators []string) int {
	acc := numbers[0]

	for i, num := range numbers[1:len(operators) + 1] {
		//fmt.Printf("Operation %d %s %d\n", acc, operators[i], num)
		switch operators[i] {
		case "*":
			acc *= num
		case "+":
			acc += num
		}
	}

	return acc
}

func extractAll(calcs []string) (testVals []int, numbers [][]int) {
	for _, line := range calcs {
		testval, numLine := extract(line)
		testVals = append(testVals, testval)
		numbers = append(numbers, numLine)
	}
	return testVals, numbers
}

func extract(calc string) (testVal int, numbers []int) {
	tokens := strings.Split(calc, ": ")
	testVal, err := strconv.Atoi(tokens[0])
	exit.If(err)
	for _, strnum := range strings.Split(tokens[1], " ") {
		numnum, err := strconv.Atoi(strnum)
		exit.If(err)
		numbers = append(numbers, numnum)
	}
	return testVal, numbers
}
