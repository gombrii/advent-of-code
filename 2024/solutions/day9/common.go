package day9

import "fmt"

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

func checksum(data []*int) int {
	sum := 0
	for i, id := range data {
		if id != nil {
			sum += i * *id
		}
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
