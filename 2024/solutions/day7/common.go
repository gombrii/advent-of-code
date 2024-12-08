package day7

import (
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
)

type manual map[string]func(int, int) int

func total(testVals []int, numbers [][]int, manual manual) int {
	acc := 0

	for i := range testVals {

		val, possible := possible(testVals[i], numbers[i], nil, manual)
		if possible {
			//fmt.Printf("%d + %d = %d\n", acc, val, acc + val)
			acc += val
		}
	}

	return acc
}

func possible(testVal int, numbers []int, operators []string, manual manual) (int, bool) {
	res := calculate(numbers, operators, manual)
	//fmt.Printf("%d=%v %v -> %d\n", testVal, numbers, operators, res)
	if res > testVal {
		return 0, false
	}
	if len(operators) >= len(numbers)-1 {
		if res == testVal {
			//fmt.Println("Found a winner!")
			return res, true
		}
		return 0, false
	}

	for operator := range manual {
		operatorAdded := append(operators[:len(operators):len(operators)], operator)
		val, poss := possible(testVal, numbers, operatorAdded, manual)
		if poss {
			return val, true
		}
	}

	return 0, false
}

func calculate(numbers []int, operators []string, manual manual) int {
	acc := numbers[0]

	for i, num := range numbers[1 : len(operators)+1] {
		//fmt.Printf("Operation %d %s %d\n", acc, operators[i], num)
		acc = manual[operators[i]](acc, num)
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
