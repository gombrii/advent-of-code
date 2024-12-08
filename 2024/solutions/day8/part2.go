package day8

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/data"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part2("2024", "day8", Part2)
}

func Part2(file string) {
	in := input.Matrix(file)

	catal := parse(in)
	antinodes := findWithHarmonics(catal, len(in[0]), len(in))

	fmt.Println(len(antinodes))
}

func findWithHarmonics(catalogue catalogue, boundX int, boundY int) map[data.Vec[int]]bool {
	antinodes := make(map[data.Vec[int]]bool)

	for _, antennae := range catalogue {
		for i, antenna := range antennae {
			for j := i + 1; j < len(antennae); j++ {
				other := antennae[j]
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
