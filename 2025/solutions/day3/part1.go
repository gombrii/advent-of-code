// Package day3 solves puzzle available on https://adventofcode.com/2025/day/3
package day3

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

type battery struct {
	joltage int
	pos     int
}

func Part1(data []byte) any {
	in := parse.Lines(data)
	banks := parseData(in)

	acc := 0
	for i, bank := range banks {
		fmt.Print("bank ", i, " ")
		acc += bankJoltage(bank)
	}

	return acc
}

func parseData(data []string) [][]battery {
	banks := make([][]battery, len(data))

	for y, bankData := range data {
		batteries := make([]battery, len(bankData))
		for x, batteryData := range bankData {
			battery, err := strconv.Atoi(string(batteryData))
			exit.If(err)

			batteries[x].joltage = battery
			batteries[x].pos = x
		}
		sort.Slice(batteries, func(i, j int) bool {
			if batteries[i].joltage != batteries[j].joltage {
				return batteries[i].joltage > batteries[j].joltage
			}
			return batteries[i].pos < batteries[j].pos
		})
		banks[y] = batteries
	}

	return banks
}

func bankJoltage(bank []battery) int {
	for _, v := range bank {
		fmt.Print(v.joltage)
	}
	fmt.Print(" ")

	first := bank[0]
	if first.pos == len(bank)-1 {
		first = bank[1]
		second := bank[0]
		fmt.Println("e", first, second)
		return first.joltage*10 + second.joltage
	}
	for _, b := range bank {
		if b == first {
			continue
		}
		second := b
		if b.pos > first.pos {
			fmt.Println(first, second)
			return first.joltage*10 + second.joltage
		}
	}

	return -1
}
