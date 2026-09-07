package decodeways

func NumDecoding(s string) int {

	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0] = 1 // empty string has exactly 1 way (do nothing)

	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}
	for i := 2; i <= n; i++ {
		oneDigit := s[i-1]      // last single digit
		twoDigits := s[i-2 : i] // last two digits as a string

		// Case 1: take the last digit alone (must not be '0')
		if oneDigit != '0' {
			dp[i] += dp[i-1]
		}

		// Case 2: take the last two digits together (must be between "10" and "26")
		if twoDigits >= "10" && twoDigits <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
