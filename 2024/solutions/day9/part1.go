package day9

import (
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registry"
)

func init() {
	registry.Register("2024", "day9", "part1", Part1)
}

func Part1(file string) any {
	in := input.String(file)

	decompressed := decompress(in)
	optimized := optimize(decompressed)
	checksum := checksum(optimized)

	return checksum
}

func appendFile(data []*int, id int, size int) []*int {
	for i := 0; i < size; i++ {
		data = append(data, &id)
	}
	return data
}

func appendSpace(data []*int, size int) []*int {
	space := make([]*int, size)
	data = append(data, space...)
	return data
}

func optimize(data []*int) (optimized []*int) {
	optimized = append(optimized, data...)

	o := 0

	for o < len(optimized) {
		for o < len(optimized) && optimized[o] != nil {
			o++
		}
		if o >= len(optimized) {
			break
		}
		optimized[o] = optimized[len(optimized)-1]
		optimized = optimized[:len(optimized)-1]
		for optimized[len(optimized)-1] == nil {
			optimized = optimized[:len(optimized)-1]
		}
		o++
	}

	return optimized
}
