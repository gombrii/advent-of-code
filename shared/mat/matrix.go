package mat

func InBounds[T any](matrix [][]T, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix[0]) && y < len(matrix)
}
