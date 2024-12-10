package day8

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/dat"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day8", "part1", Part1)
}

type catalogue map[string][]dat.Vec[int]

func Part1(file string) {
	in := input.Matrix(file)

	catal := parse(in)
	antinodes := findAntinodes(catal, len(in[0]), len(in))

	fmt.Println(len(antinodes))
}

func findAntinodes(catalogue catalogue, boundX int, boundY int) map[dat.Vec[int]]bool {
	antinodes := make(map[dat.Vec[int]]bool)

	for _, antennae := range catalogue {
		for i, antenna := range antennae {
			for _, other := range antennae[i+1:] {
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
