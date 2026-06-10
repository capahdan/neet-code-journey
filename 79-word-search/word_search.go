package wordsearch

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {

		if idx == len(word) {
			return true
		}

		if r < 0 || r >= m || c < 0 || c >= n {
			return false
		}

		if board[r][c] != word[idx] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		found := dfs(r+1, c, idx+1) || dfs(r, c+1, idx+1) || dfs(r-1, c, idx+1) || dfs(r, c-1, idx+1)
		board[r][c] = temp

		return found
	}

	for r := range board {
		for c := range board[r] {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}
