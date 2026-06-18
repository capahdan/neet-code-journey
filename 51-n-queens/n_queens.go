package nqueens

func solveNQueens(n int) [][]string {
	col := make(map[int]bool)
	posDiag := make(map[int]bool) // (r + c)
	negDiag := make(map[int]bool) // (r - c)

	var res [][]string

	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var backTrack func(r int)
	backTrack = func(r int) {
		if r == n {
			var boardCopy []string
			for i := 0; i < n; i++ {
				boardCopy = append(boardCopy, string(board[i]))
			}
			res = append(res, boardCopy)
			return
		}

		// Try placing a queen in each column of the current row
		for c := 0; c < n; c++ {
			// Check if the current square is already under attack
			if col[c] || posDiag[r+c] || negDiag[r-c] {
				continue
			}

			// Place the queen and mark the column and diagonals as occupied
			col[c] = true
			posDiag[r+c] = true
			negDiag[r-c] = true
			board[r][c] = 'Q'

			// Recursively move to the next row
			backTrack(r + 1)

			// Backtrack: Clean up the board and sets for the next iteration
			col[c] = false
			posDiag[r+c] = false
			negDiag[r-c] = false
			board[r][c] = '.'
		}
	}

	return res
}

// the reason why that solution works because we are implementing backtracking
// backtracking algoritm, we try to do all posibiilities in a function

// let's say we have n is 4

// . Q . .
// . . . Q
// Q . . .
// . . Q .

// why that pattern work because in the chess the queen can kill each other
// when they are in the same colum and
