package day8

import (
	"github.com/gombrii/Advent-of-code/shared/dat"
	parse1 "github.com/gombrii/Advent-of-code/shared/parse"
)

func Part1(data []byte) any {
	in := parse1.Matrix(data, "")

	catal := parse(in)
	antinodes := findAntinodes(catal, len(in[0]), len(in))

	return len(antinodes)
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
