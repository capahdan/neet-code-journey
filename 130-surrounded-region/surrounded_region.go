package surroundedregion

// the problem statement
// surrounded region is those region of 'O' which are not connected to the border
// it is in the 'X' matrix it adjectence is only to top, bottom, right, left
// when the region is surrounded with 'X' we need to change it to O but when the
// when the region is in the border we don't need to change it
// so the first thing we sould do is to transverse all the border and change it to
// temporary caracter which is # and then we check all the value each one
// when we found character of # we restore to O but when we find character of 'O'
// we flip it to 'X'

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = '#' // safe: connected to the border
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	// Phase 1: start from every border cell
	for i := 0; i < rows; i++ {
		dfs(i, 0)
		dfs(i, cols-1)
	}
	for j := 0; j < cols; j++ {
		dfs(0, j)
		dfs(rows-1, j)
	}

	// Phases 2 and 3: flip surrounded O to X, restore # to O
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case '#':
				board[i][j] = 'O'
			}
		}
	}

}
