package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := range n / 2 {
		for j := i; j < n-i-1; j++ {
			top := matrix[i][j]
			right := matrix[j][n-i-1]
			bottom := matrix[n-i-1][n-j-1]
			left := matrix[n-j-1][i]

			matrix[i][j] = left
			matrix[j][n-i-1] = top
			matrix[n-i-1][n-j-1] = right
			matrix[n-j-1][i] = bottom
		}
	}
}
