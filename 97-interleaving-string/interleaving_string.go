package interleavingstring

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}

	// dp[i][j] = can s1[i:] and s2[j:] build s3[i+j:] ?
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[n][m] = true // nothing left in either string: empty matches empty

	for i := n; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			if i < n && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
			}
			if j < m && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	return dp[0][0]
}
