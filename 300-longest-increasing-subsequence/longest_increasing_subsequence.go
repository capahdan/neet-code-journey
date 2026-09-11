package longestincreasingsubsequence

func LengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range nums {
		dp[i] = 1
	}

	longest := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > longest {
			longest = dp[i]
		}
	}

	return longest
}
