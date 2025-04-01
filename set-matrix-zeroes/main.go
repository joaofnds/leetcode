package main

func setZeroes(matrix [][]int) {
	rowsToZero := make([]bool, len(matrix))
	colsToZero := make([]bool, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				rowsToZero[i] = true
				colsToZero[j] = true
			}
		}
	}

	for i := range matrix {
		for j := range matrix[0] {
			if rowsToZero[i] || colsToZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}
