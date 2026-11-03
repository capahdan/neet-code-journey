package editdistance

func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= len(word1); i++ {
		dp[i][len(word2)] = len(word1) - i
	}

	for j := 0; j <= len(word2); j++ {
		dp[len(word1)][j] = len(word2) - j
	}

	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j+1], dp[i+1][j+1])
			}
		}
	}

	return dp[0][0]
}
