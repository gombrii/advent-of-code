package matrices

import "fmt"

func Println(matrix [][]string) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			fmt.Printf("[%s]", matrix[y][x])
		}
		if y != len(matrix)-1 {
			fmt.Println()
		}
	}
}
