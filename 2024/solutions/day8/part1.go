package day8

import (
	"fmt"
	"regexp"

	"github.com/gomsim/Advent-of-code/shared/data"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day8", Part1)
}

var antenna = regexp.MustCompile(`\w|\d`)

type catalogue map[string][]data.Vec[int]

func Part1(file string) {
	in := input.Matrix(file)

	catal := parse(in)
	antinodes := findAntinodes(catal, len(in[0]), len(in))

	fmt.Println(len(antinodes))
}

func parse(matrix [][]string) catalogue {
	catal := make(catalogue)

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			marking := matrix[y][x]
			if antenna.MatchString(marking) {
				catal[marking] = append(catal[marking], data.Vec[int]{X: x, Y: y})
			}
		}
	}

	return catal
}

func findAntinodes(catalogue catalogue, boundX int, boundY int) map[data.Vec[int]]bool {
	antinodes := make(map[data.Vec[int]]bool)

	for _, antennae := range catalogue {
		for i, antenna := range antennae {
			for j := i + 1; j < len(antennae); j++ {
				other := antennae[j]
				dist := antenna.Sub(other)
				antiA := antenna.Add(dist)
				if antiA.X >= 0 && antiA.Y >= 0 && antiA.X < boundX && antiA.Y < boundY {
					antinodes[antiA] = true
				}
				antiB := other.Sub(dist)
				if antiB.X >= 0 && antiB.Y >= 0 && antiB.X < boundX && antiB.Y < boundY {
					antinodes[antiB] = true
				}
			}
		}
	}

	return antinodes
}
