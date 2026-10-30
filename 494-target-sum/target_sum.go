package targetsum

func TargetSum(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1

	for _, n := range nums {
		next := make(map[int]int)
		for sum, ways := range dp {
			next[sum+n] += ways
			next[sum-n] += ways
		}
		dp = next
	}

	return dp[target]
}
