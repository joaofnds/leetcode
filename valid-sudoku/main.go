package main

func isValidSudoku(board [][]byte) bool {
	for row := range 9 {
		if !isValid(board[row]) {
			return false
		}
	}

	for col := range 9 {
		column := make([]byte, 0, 9)
		for row := range 9 {
			column = append(column, board[row][col])
		}
		if !isValid(column) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			subgrid := make([]byte, 0, 9)
			for k := range 3 {
				for l := range 3 {
					subgrid = append(subgrid, board[i+k][j+l])
				}
			}
			if !isValid(subgrid) {
				return false
			}
		}
	}

	return true
}

func isValid(set []byte) bool {
	seen := map[byte]bool{}
	for _, b := range set {
		if b == '.' {
			continue
		}
		if seen[b] {
			return false
		}
		seen[b] = true
	}
	return true
}
