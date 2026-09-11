package wordbreak

func WordBreak(s string, wordDict []string) bool {

	wordSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}

	}

	return dp[n]
}
