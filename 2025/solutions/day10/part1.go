// Package day10 solves puzzle available on https://adventofcode.com/2025/day/10
package day10

import (
	"bytes"
	"slices"
	"strconv"
	"strings"

	"github.com/gombrii/Advent-of-code/shared/exit"
	"github.com/gombrii/Advent-of-code/shared/parse"
)

type crawler struct {
	light uint16
	path  []uint16
}

func Part1(data []byte) any {
	data = bytes.ReplaceAll(data, []byte("#"), []byte("1"))
	data = bytes.ReplaceAll(data, []byte("."), []byte("0"))
	in := parse.Lines(data)
	lArrs, bArrs, jReqs := compile(in)

	acc := 0
	for i := range len(in) {
		//fmt.Println("MACHINE", i)
		//fmt.Println(render.MapSlice(bArrs[i], func(a uint16) any { return fmt.Sprintf("%016b", a) }))
		acc += leastPresses(lArrs[i], bArrs[i], jReqs[i])
	}

	return acc
}

func leastPresses(light uint16, buttons []uint16, joltReq []int) int {
	crawlers := []crawler{}
	for _, button := range buttons {
		crawlers = append(crawlers, crawler{light: light, path: []uint16{button}})
	}

	for {
		newCrawlers := []crawler{}

		for _, c := range crawlers {
			tip := c.path[len(c.path)-1]
			if c.light ^= tip; c.light == 0 {
				return len(c.path)
			}

			for _, button := range buttons {
				if !slices.Contains(c.path, button) {
					newCrawlers = append(newCrawlers, crawler{light: c.light, path: append(append([]uint16{}, c.path...), button)})
				}
			}
		}

		crawlers = append(crawlers, newCrawlers...)
	}
}

func compile(data []string) (lArrs []uint16, bArrs [][]uint16, jReqs [][]int) {
	lArrs = make([]uint16, len(data))
	bArrs = make([][]uint16, len(data))
	jReqs = make([][]int, len(data))

	for l, line := range data {
		parts := strings.Split(line, " ")
		buttons := make([]uint16, len(parts)-2)
		for p, part := range parts {
			switch part[0] {
			case '[':
				part = strings.TrimSuffix(strings.TrimPrefix(part, "["), "]")
				bString, err := strconv.ParseUint(part, 2, 16)
				exit.If(err)
				lArrs[l] = uint16(bString)
			case '(':
				button := uint16(0)
				nums := strings.TrimSuffix(strings.TrimPrefix(part, "("), ")")
				for _, num := range strings.Split(nums, ",") {
					n, err := strconv.Atoi(num)
					exit.If(err)
					button |= 1 << (len(parts[0]) - 2 - n - 1)
				}
				buttons[p-1] = button
			case '{':
				jRec := []int{}
				nums := strings.TrimSuffix(strings.TrimPrefix(part, "{"), "}")
				for _, num := range strings.Split(nums, ",") {
					n, err := strconv.Atoi(num)
					exit.If(err)
					jRec = append(jRec, n)
				}
				jReqs[l] = jRec
			}
		}
		bArrs[l] = buttons
	}

	return lArrs, bArrs, jReqs
}
