package uniquepath

// we are using dfs with memoization
// func UniquePath(m, n int) int {
// 	memo := make([][]int, m)
// 	for i := range memo {
// 		memo[i] = make([]int, n)
// 	}

// 	var dfs func(r, c int) int
// 	dfs = func(r, c int) int {
// 		if r >= m || c >= n {
// 			return 0
// 		}

// 		if r == m-1 && c == n-1 {
// 			return 1
// 		}

// 		if memo[r][c] != 0 {
// 			return memo[r][c]
// 		}

// 		memo[r][c] = dfs(r+1, c) + dfs(r, c+1)
// 		return memo[r][c]
// 	}

// 	return dfs(0, 0)
// }

func UniquePath(m, n int) int {

	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}

// func UniquePath(m int, n int) int {
// 	N := m + n - 2
// 	k := min(m-1, n-1)
// 	res := 1
// 	for i := 1; i <= k; i++ {
// 		res = res * (N - k + i) / i
// 	}
// 	return res
// }
