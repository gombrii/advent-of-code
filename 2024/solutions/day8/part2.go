package day8

import (
	"github.com/gomsim/Advent-of-code/shared/dat"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day8", "part2", Part2)
}

func Part2(file string) any {
	in := input.Matrix(file)

	catal := parse(in)
	antinodes := findWithHarmonics(catal, len(in[0]), len(in))

	return len(antinodes)
}

func findWithHarmonics(catalogue catalogue, boundX int, boundY int) map[dat.Vec[int]]bool {
	antinodes := make(map[dat.Vec[int]]bool)

	for _, antennae := range catalogue {
		for i, antenna := range antennae {
			for _, other := range antennae[i+1:] {
				dir := antenna.Sub(other).Red()

				anti := antenna
				for anti.X >= 0 && anti.Y >= 0 && anti.X < boundX && anti.Y < boundY {
					antinodes[anti] = true
					anti = anti.Add(dir)
				}
				anti = antenna
				for anti.X >= 0 && anti.Y >= 0 && anti.X < boundX && anti.Y < boundY {
					antinodes[anti] = true
					anti = anti.Sub(dir)
				}
			}
		}
	}

	return antinodes
}
