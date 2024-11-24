package main

import (
	"fmt"
	"os"

	"github.com/gomsim/Advent-of-code/2023/solutions/day1"
	"github.com/gomsim/Advent-of-code/2023/solutions/day2"
	"github.com/gomsim/Advent-of-code/shared/exit"
)

const (
	year  = 2023

	dayIndex   = 1
	partIndex  = 2
	inputIndex = 3
)

func main() {
	if len(os.Args) != 4 {
		exit.Errorf("Need 3 arguments (daynumber, partnumber, input filename), got %d", len(os.Args)-1)
	}

	day := fmt.Sprintf("day%s", os.Args[dayIndex])
	part := fmt.Sprintf("part%s", os.Args[partIndex])
	inputFile := fmt.Sprintf("%d/input/%s/%s.txt", year, day, os.Args[inputIndex])

	navigate := navigationalMap()

	fmt.Printf("Running %d/%s/%s with %s\n", year, day, part, inputFile)

	navigate[day][part](inputFile)
}

func navigationalMap() map[string]map[string]func(string){
	return map[string]map[string]func(string){
		"day1": {
			"part1": func(inputFile string){day1.Part1(inputFile)},
			"part2": func(inputFile string){day1.Part2(inputFile)},
		},
		"day2": {
			"part1": func(inputFile string){day2.Part1(inputFile)},
			"part2": func(inputFile string){day2.Part2(inputFile)},
		},
		//"day3": {
		//	"part1": func(inputFile string){day3.Part1(inputFile)},
		//	"part2": func(inputFile string){day3.Part2(inputFile)},
		//},
		//"day4": {
		//	"part1": func(inputFile string){day4.Part1(inputFile)},
		//	"part2": func(inputFile string){day4.Part2(inputFile)},
		//},
		//"day5": {
		//	"part1": func(inputFile string){day5.Part1(inputFile)},
		//	"part2": func(inputFile string){day5.Part2(inputFile)},
		//},
		//"day6": {
		//	"part1": func(inputFile string){day6.Part1(inputFile)},
		//	"part2": func(inputFile string){day6.Part2(inputFile)},
		//},
		//"day7": {
		//	"part1": func(inputFile string){day7.Part1(inputFile)},
		//	"part2": func(inputFile string){day7.Part2(inputFile)},
		//},
		//"day8": {
		//	"part1": func(inputFile string){day8.Part1(inputFile)},
		//	"part2": func(inputFile string){day8.Part2(inputFile)},
		//},
		//"day9": {
		//	"part1": func(inputFile string){day9.Part1(inputFile)},
		//	"part2": func(inputFile string){day9.Part2(inputFile)},
		//},
		//"day10": {
		//	"part1": func(inputFile string){day10.Part1(inputFile)},
		//	"part2": func(inputFile string){day10.Part2(inputFile)},
		//},
		//"day11": {
		//	"part1": func(inputFile string){day11.Part1(inputFile)},
		//	"part2": func(inputFile string){day11.Part2(inputFile)},
		//},
		//"day12": {
		//	"part1": func(inputFile string){day12.Part1(inputFile)},
		//	"part2": func(inputFile string){day12.Part2(inputFile)},
		//},
		//"day13": {
		//	"part1": func(inputFile string){day13.Part1(inputFile)},
		//	"part2": func(inputFile string){day13.Part2(inputFile)},
		//},
		//"day14": {
		//	"part1": func(inputFile string){day14.Part1(inputFile)},
		//	"part2": func(inputFile string){day14.Part2(inputFile)},
		//},
		//"day15": {
		//	"part1": func(inputFile string){day15.Part1(inputFile)},
		//	"part2": func(inputFile string){day15.Part2(inputFile)},
		//},
		//"day16": {
		//	"part1": func(inputFile string){day16.Part1(inputFile)},
		//	"part2": func(inputFile string){day16.Part2(inputFile)},
		//},
		//"day17": {
		//	"part1": func(inputFile string){day17.Part1(inputFile)},
		//	"part2": func(inputFile string){day17.Part2(inputFile)},
		//},
		//"day18": {
		//	"part1": func(inputFile string){day18.Part1(inputFile)},
		//	"part2": func(inputFile string){day18.Part2(inputFile)},
		//},
		//"day19": {
		//	"part1": func(inputFile string){day19.Part1(inputFile)},
		//	"part2": func(inputFile string){day19.Part2(inputFile)},
		//},
		//"day20": {
		//	"part1": func(inputFile string){day20.Part1(inputFile)},
		//	"part2": func(inputFile string){day20.Part2(inputFile)},
		//},
		//"day21": {
		//	"part1": func(inputFile string){day21.Part1(inputFile)},
		//	"part2": func(inputFile string){day21.Part2(inputFile)},
		//},
		//"day22": {
		//	"part1": func(inputFile string){day22.Part1(inputFile)},
		//	"part2": func(inputFile string){day22.Part2(inputFile)},
		//},
		//"day23": {
		//	"part1": func(inputFile string){day23.Part1(inputFile)},
		//	"part2": func(inputFile string){day23.Part2(inputFile)},
		//},
		//"day24": {
		//	"part1": func(inputFile string){day24.Part1(inputFile)},
		//	"part2": func(inputFile string){day24.Part2(inputFile)},
		//},
		//"day25": {
		//	"part1": func(inputFile string){day25.Part1(inputFile)},
		//	"part2": func(inputFile string){day25.Part2(inputFile)},
		//},
	}
}
