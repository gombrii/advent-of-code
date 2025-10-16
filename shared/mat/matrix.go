package mat

func InMatrix[T any](matrix [][]T, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(matrix[0]) && y < len(matrix)
}

func InSlice[T any](slice []T, i int) bool {
	return i >= 0 && i < len(slice)
}
