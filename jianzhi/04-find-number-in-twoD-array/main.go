package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5

	// matrix := [][]int{{1, 1}}
	// target := 2
	fmt.Println(findNumberIn2DArray(matrix, target))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	x := 0
	y := len(matrix) - 1

	width := len(matrix[0])

	for x < width && y >= 0 {
		if matrix[y][x] == target {
			return true
		}

		if matrix[y][x] > target {
			y--
		} else if matrix[y][x] < target {
			x++
		}
	}

	return false
}
