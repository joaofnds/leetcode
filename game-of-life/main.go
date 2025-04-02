package main

func gameOfLife(board [][]int) {
	lookup := func(i, j int) int {
		if 0 <= i && i < len(board) && 0 <= j && j < len(board[i]) {
			return board[i][j] & 1
		}
		return 0
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] += (lookup(i-1, j-1) + lookup(i-1, j) + lookup(i-1, j+1) +
				lookup(i, j-1) + lookup(i, j+1) +
				lookup(i+1, j-1) + lookup(i+1, j) + lookup(i+1, j+1)) << 1
		}
	}

	for i := range board {
		for j := range board[i] {
			cell, neighbors := board[i][j]&1, board[i][j]>>1

			if cell == 1 && (neighbors < 2 || neighbors > 3) {
				board[i][j] = 0
			} else if cell == 0 && neighbors == 3 {
				board[i][j] = 1
			} else {
				board[i][j] = cell
			}
		}
	}
}
