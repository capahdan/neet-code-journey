package pasificatlanticwaterflow

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])

	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// dfs floods "uphill" from an ocean border cell, marking every cell that
	// can drain into that ocean (i.e. heights[nr][nc] >= heights[r][c]).
	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true
		dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			if visited[nr][nc] {
				continue
			}
			if heights[nr][nc] < heights[r][c] {
				continue
			}
			dfs(nr, nc, visited)
		}
	}

	for c := 0; c < n; c++ {
		dfs(0, c, pacific)    // top row touches Pacific
		dfs(m-1, c, atlantic) // bottom row touches Atlantic
	}
	for r := 0; r < m; r++ {
		dfs(r, 0, pacific)    // left column touches Pacific
		dfs(r, n-1, atlantic) // right column touches Atlantic
	}

	result := [][]int{}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
