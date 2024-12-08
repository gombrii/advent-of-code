package day8

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/data"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part1("2024", "day8", Part1)
}

type catalogue map[string][]data.Vec[int]

func Part1(file string) {
	in := input.Matrix(file)

	catal := parse(in)
	antinodes := findAntinodes(catal, len(in[0]), len(in))

	fmt.Println(len(antinodes))
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
