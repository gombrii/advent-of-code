package day9

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day9", "part1", Part1)
}

func Part1(file string) {
	in := input.String(file)

	//fmt.Println(in)
	decompressed := decompress(in)
	//PrintData(decompressed)
	optimized := optimize(decompressed)
	//PrintData(optimized)
	checksum := checksum(optimized)

	fmt.Println(checksum)
}

func decompress(compressed string) []*int {
	data := make([]*int, 0, len(compressed)*2)
	count := 0
	for i, val := range compressed {
		size := int(val - '0')
		if i%2 == 0 {
			data = appendFile(data, count, size)
			count++
		} else {
			data = appendSpace(data, size)
		}
	}
	return data
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
		//fmt.Print("o:", o, " -> ")
		//PrintData(optimized)
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

func checksum(data []*int) int {
	sum := 0
	for i, id := range data {
		sum += i * *id
	}
	return sum
}

func PrintData(data []*int) {
	for _, block := range data {
		if block == nil {
			fmt.Print(".")
		} else {
			fmt.Print(*block)
		}
	}
	fmt.Println()
}
