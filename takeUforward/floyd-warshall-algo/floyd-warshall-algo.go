package main

import "fmt"

const INT_MAX int = 1 << 32

func floyd_warshall(matrix [][]int) [][]int {
	N := len(matrix)
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				matrix[i][j] = min(matrix[i][j], matrix[i][k]+matrix[k][j])
			}
		}
	}

	fmt.Println(matrix)
	return matrix
}

func main() {
	matrix := [][]int{{0, 2, -1, -1},
		{1, 0, 3, -1},
		{-1, -1, 0, 1},
		{3, 5, 4, 0},
	}

	N := len(matrix)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = INT_MAX
			}
		}
	}

	floyd_warshall(matrix)
}
